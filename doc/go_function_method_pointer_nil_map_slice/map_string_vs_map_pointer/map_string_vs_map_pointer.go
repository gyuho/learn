package map_string_vs_map_pointer

import "sync"

type Interface interface {
	set(v *node)
	exist(id string) bool
	delete(id string)
}

type node struct {
	id string
}

type MapString struct {
	mu   sync.Mutex
	data map[string]*node
}

func newMapString() *MapString {
	d := &MapString{}
	d.data = make(map[string]*node)
	return d
}

func (d *MapString) set(v *node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v.id] = v
}

func (d *MapString) exist(id string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, ok := d.data[id]
	return ok
}

func (d *MapString) delete(id string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[id]; ok {
		delete(d.data, id)
	}
}

type MapPointer struct {
	mu   sync.Mutex
	data map[*node]struct{}
}

func newMapPointer() *MapPointer {
	d := &MapPointer{}
	d.data = make(map[*node]struct{})
	return d
}

func (d *MapPointer) set(v *node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v] = struct{}{}
}

func (d *MapPointer) exist(id string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	for v := range d.data {
		if v.id == id {
			return true
		}
	}
	return false
}

func (d *MapPointer) delete(id string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for v := range d.data {
		if v.id == id {
			delete(d.data, v)
			return
		}
	}
}
