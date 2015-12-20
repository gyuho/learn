package slice_vs_map

import "sync"

type Interface interface {
	set(v string)
	exist(v string) bool
	delete(v string)
}

type Slice struct {
	mu   sync.Mutex
	data []string
}

type Map struct {
	mu   sync.Mutex
	data map[string]struct{}
}

func newSlice() *Slice {
	d := &Slice{}
	d.data = make([]string, 0)
	return d
}

func (d *Slice) set(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data = append(d.data, v)
}

func (d *Slice) exist(v string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, vv := range d.data {
		if vv == v {
			return true
		}
	}
	return false
}

func (d *Slice) delete(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	del := -1
	for i, vv := range d.data {
		if vv == v {
			del = i
		}
	}
	if del != -1 {
		d.data = append(d.data[:del], d.data[del+1:]...)
	}
}

func newMap() *Map {
	d := &Map{}
	d.data = make(map[string]struct{})
	return d
}

func (d *Map) set(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v] = struct{}{}
}

func (d *Map) exist(v string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, ok := d.data[v]
	return ok
}

func (d *Map) delete(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[v]; ok {
		delete(d.data, v)
	}
}
