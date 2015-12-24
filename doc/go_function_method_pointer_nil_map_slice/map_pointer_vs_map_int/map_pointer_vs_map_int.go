package map_pointer_vs_map_int

import "sync"

type Interface interface {
	set(v *metric)
	exist(v *metric) bool
	delete(v *metric)
}

type metric struct {
	id int
	s  string
}

func newMapPointer() *MapPointer {
	d := &MapPointer{}
	d.data = make(map[*metric]struct{})
	return d
}

func (d *MapPointer) set(v *metric) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v] = struct{}{}
}

func (d *MapPointer) exist(v *metric) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, ok := d.data[v]
	return ok
}

func (d *MapPointer) delete(v *metric) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[v]; ok {
		delete(d.data, v)
	}
}

type MapInt struct {
	mu   sync.Mutex
	data map[int]*metric
}

type MapPointer struct {
	mu   sync.Mutex
	data map[*metric]struct{}
}

func newMapInt() *MapInt {
	d := &MapInt{}
	d.data = make(map[int]*metric)
	return d
}

func (d *MapInt) set(v *metric) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v.id] = v
}

func (d *MapInt) exist(v *metric) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, ok := d.data[v.id]
	return ok
}

func (d *MapInt) delete(v *metric) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[v.id]; ok {
		delete(d.data, v.id)
	}
}
