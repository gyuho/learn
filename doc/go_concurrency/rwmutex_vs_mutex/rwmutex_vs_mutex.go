package rwmutex_vs_mutex

import "sync"

type Interface interface {
	set(v string)
	exist(v string) bool
	delete(v string)
}

type DataRWMutex struct {
	mu   sync.RWMutex
	data map[string]struct{}
}

type DataMutex struct {
	mu   sync.Mutex
	data map[string]struct{}
}

func newDataRWMutex() *DataRWMutex {
	d := &DataRWMutex{}
	d.data = make(map[string]struct{}, 0)
	return d
}

func (d *DataRWMutex) set(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v] = struct{}{}
}

func (d *DataRWMutex) exist(v string) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	_, ok := d.data[v]
	return ok
}

func (d *DataRWMutex) delete(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[v]; ok {
		delete(d.data, v)
	}
}

func newDataMutex() *DataMutex {
	d := &DataMutex{}
	d.data = make(map[string]struct{})
	return d
}

func (d *DataMutex) set(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[v] = struct{}{}
}

func (d *DataMutex) exist(v string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, ok := d.data[v]
	return ok
}

func (d *DataMutex) delete(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.data[v]; ok {
		delete(d.data, v)
	}
}
