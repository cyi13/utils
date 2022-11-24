package funcs

type Top struct {
	cap  int
	data []*TopValue
}

type TopValue struct {
	Key string
	Num int
}

func NewTop(cap int) *Top {
	return &Top{
		data: make([]*TopValue, 0, cap),
		cap:  cap,
	}
}

func (t *Top) Add(key string, num int) {
	for i, v := range t.data {
		if v.Num < num {
			t.data = append(t.data[:i], append([]*TopValue{{Key: key, Num: num}}, t.data[i:]...)...)
			t.checkCap()
			return
		}
	}
	t.data = append(t.data, &TopValue{Key: key, Num: num})
	t.checkCap()
}

func (t *Top) checkCap() {
	if len(t.data) > t.cap {
		t.data = t.data[:t.cap]
	}
}

func (t *Top) Result() []*TopValue {
	return t.data
}
