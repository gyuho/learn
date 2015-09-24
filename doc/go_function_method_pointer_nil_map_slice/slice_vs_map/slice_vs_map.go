package slice_vs_map

import "sync"

type message struct {
	body string
}

func newMessage(body string) *message {
	m := new(message)
	m.body = body
	return m
}

type Interface interface {
	Add(msg *message)
	Exist(msg *message) bool
	Delete(msg *message)
}

type Slice struct {
	sync.Mutex
	data []*message
}

func newSlice() *Slice {
	d := new(Slice)
	d.data = make([]*message, 0)
	return d
}

func (s *Slice) Add(msg *message) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, msg)
}

func (s *Slice) Exist(msg *message) bool {
	s.Lock()
	defer s.Unlock()
	for _, v := range s.data {
		if v == msg {
			return true
		}
	}
	return false
}

func (s *Slice) Delete(msg *message) {
	s.Lock()
	defer s.Unlock()
	del := -1
	for i, v := range s.data {
		if v == msg {
			del = i
		}
	}
	s.data = append(s.data[:del], s.data[del+1:]...)
}

type Map struct {
	sync.Mutex
	data map[*message]struct{}
}

func newMap() *Map {
	d := new(Map)
	d.data = make(map[*message]struct{})
	return d
}

func (m *Map) Add(msg *message) {
	m.Lock()
	defer m.Unlock()
	m.data[msg] = struct{}{}
}

func (m *Map) Exist(msg *message) bool {
	m.Lock()
	defer m.Unlock()
	_, ok := m.data[msg]
	return ok
}

func (m *Map) Delete(msg *message) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, msg)
}
