package map_pointer_vs_map_int

import "sync"

type Interface interface {
	set(v *node)
	exist(v *node) bool
	delete(v *node)
}

type node struct {
	id int
	s  string
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

func (d *MapPointer) exist(v *node) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, ok := d.data[v]
	return ok
}

func (d *MapPointer) delete(v *node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[v]; ok {
		delete(d.data, v)
	}
}

type MapInt struct {
	mu   sync.Mutex
	data map[int]*node
}

func newMapInt() *MapInt {
	d := &MapInt{}
	d.data = make(map[int]*node)
	return d
}

func (d *MapInt) set(v *node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v.id] = v
}

func (d *MapInt) exist(v *node) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, ok := d.data[v.id]
	return ok
}

func (d *MapInt) delete(v *node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[v.id]; ok {
		delete(d.data, v.id)
	}
}
