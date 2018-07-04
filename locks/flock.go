package locks

import (
	"os"
	"syscall"
)

//fileLock 实现一个队单个文件加锁的封装
type fileLock struct {
	f *os.File
}

//NewFlock fileLock
func NewFlock(dir string) (*fileLock, error) {
	f, err := os.Open(dir)
	return &fileLock{
		f: f,
	}, err
}

//IsLock 文件排它锁 当检测到文件被上锁的时候 会返回错误 也可以用来检测文件是否上被锁
func (l *fileLock) IsLock() error {
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

//Lock 文件排它锁 当文件被上锁的时候 会阻塞等待锁的释放
func (l *fileLock) Lock() {
	syscall.Flock(int(l.f.Fd()), syscall.LOCK_EX)
}

//UnLock 解除锁
func (l *fileLock) UnLock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
