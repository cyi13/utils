package fsnotify

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	watcher  *fsnotify.Watcher
	wathPath string

	skipDir   []string
	dirEvent  chan fsnotify.Event
	watchList sync.Map

	out    chan *Message
	custom chan fsnotify.Event
	done   chan struct{}
}

const (
	EVENT_UPDATE = iota
	EVENT_DELETE
)

type Message struct {
	File  string
	Event int
}

func NewWatch(path string, skipDir []string) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	w := &Watcher{
		skipDir:  skipDir,
		watcher:  watcher,
		wathPath: path,
		dirEvent: make(chan fsnotify.Event, 100),
		out:      make(chan *Message, 100),
		done:     make(chan struct{}),
		custom:   make(chan fsnotify.Event, 10),
	}
	if err := w.init(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Watcher) Listen() chan *Message {
	return w.out
}

func (w *Watcher) Close() {
	close(w.done)
}

func (w *Watcher) init() error {
	if err := filepath.WalkDir(w.wathPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}

		if w.isSkipDir(path) {
			return fs.SkipDir
		}

		return w.Add(path)
	}); err != nil {
		return err
	}

	go w.run()
	go w.watchNewDir()
	return nil
}

func (w *Watcher) run() {
	ticker := time.NewTicker(5 * time.Second)
	var pendding []fsnotify.Event
	for {
		select {
		case <-ticker.C:
			w.processFile(pendding)
			pendding = []fsnotify.Event{}
		case <-w.done:
			w.processFile(pendding)
			close(w.out)
			w.watcher.Close()
		case v := <-w.watcher.Events:
			//新的目录或者在监视列表
			if w.isDir(v.Name) || w.isWatching(v.Name) {
				w.dirEvent <- v
				continue
			}
			pendding = append(pendding, v)
			//防止积压过多数据
			if len(pendding) > 1000 {
				w.processFile(pendding)
				pendding = []fsnotify.Event{}
			}
		case v := <-w.custom:
			pendding = append(pendding, v)
			if len(pendding) > 1000 {
				w.processFile(pendding)
				pendding = []fsnotify.Event{}
			}
		}
	}
}

// 文件批量处理，批量是为了处理一个文件同时有多个事件的问题，比如改名等
func (w *Watcher) processFile(data []fsnotify.Event) {
	files := make(map[string]fsnotify.Event)
	for _, v := range data {
		if w.chmodEvent(v.Op) || w.renameEvent(v.Op) {
			continue
		}
		files[v.Name] = v
	}

	for _, v := range files {
		msg := w.event2Msg(v)
		if msg == nil {
			continue
		}

		w.out <- msg
	}
}

// 文件事件转换成消息
func (w *Watcher) event2Msg(event fsnotify.Event) *Message {
	switch {
	case w.createEvent(event.Op), w.writeEvent(event.Op):
		return &Message{
			Event: EVENT_UPDATE,
			File:  event.Name,
		}
	case w.removeEvent(event.Op):
		return &Message{
			Event: EVENT_DELETE,
			File:  event.Name,
		}
	}

	return nil
}

func (w *Watcher) watchNewDir() {
	for {
		select {
		case <-w.done:
			return
		case event := <-w.dirEvent:
			w.processDir(event)
		}
	}
}

func (w *Watcher) processDir(data fsnotify.Event) {
	for _, v := range w.skipDir {
		if strings.HasSuffix(data.Name, v) {
			return
		}
	}
	switch {
	case w.createEvent(data.Op):
		w.processNewDir(data.Name)
	case w.removeEvent(data.Op):
		w.processDeleteDir(data.Name)
	}
}

func (w *Watcher) processNewDir(dir string) {
	if w.isWatching(dir) {
		return
	}
	if w.isSkipDir(dir) {
		return
	}
	w.Add(dir)
	entrys, _ := os.ReadDir(dir)
	for _, v := range entrys {
		path := filepath.Join(dir, v.Name())
		if v.IsDir() {
			w.processNewDir(path)
			continue
		}
		w.custom <- fsnotify.Event{
			Name: path,
			Op:   fsnotify.Create,
		}
	}
}

func (w *Watcher) processDeleteDir(dir string) {
	if err := w.Remove(dir); err != nil {
	}
}

func (w *Watcher) isWatching(name string) bool {
	_, ok := w.watchList.Load(name)
	return ok
}

func (w *Watcher) isSkipDir(name string) bool {
	for _, v := range w.skipDir {
		if strings.HasSuffix(name, v) {
			return true
		}
	}
	return false
}

func (w *Watcher) Add(name string) error {
	if w.isWatching(name) {
		return nil
	}
	if err := w.watcher.Add(name); err != nil {
		return err
	}

	w.watchList.Store(name, struct{}{})
	return nil
}

func (w *Watcher) Remove(name string) error {
	w.watchList.Delete(name)
	if err := w.watcher.Remove(name); err != nil {
		return err
	}
	return nil
}

func (w *Watcher) Stat() int {
	count := 0
	w.watchList.Range(func(_, _ interface{}) bool {
		count++
		return true
	})

	return count
}

func (w *Watcher) writeEvent(op fsnotify.Op) bool {
	return op&fsnotify.Write == fsnotify.Write
}

func (w *Watcher) createEvent(op fsnotify.Op) bool {
	return op&fsnotify.Create == fsnotify.Create
}

func (w *Watcher) removeEvent(op fsnotify.Op) bool {
	return op&fsnotify.Remove == fsnotify.Remove
}

func (w *Watcher) chmodEvent(op fsnotify.Op) bool {
	return op&fsnotify.Chmod == fsnotify.Chmod
}

func (w *Watcher) renameEvent(op fsnotify.Op) bool {
	return op&fsnotify.Rename == fsnotify.Rename
}

func (w *Watcher) isDir(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return info.IsDir()
}
