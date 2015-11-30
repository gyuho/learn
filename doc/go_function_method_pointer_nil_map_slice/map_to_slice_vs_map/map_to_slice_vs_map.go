package map_to_slice_vs_map

import "sync"

type Interface interface {
	set(k, v string)
	exist(k, v string) bool
	delete(k, v string)
}

type mapToSlice struct {
	mu   sync.Mutex
	data map[string][]string
}

type mapToMap struct {
	mu   sync.Mutex
	data map[string]map[string]struct{}
}

func newMapToSlice() *mapToSlice {
	d := &mapToSlice{}
	d.data = make(map[string][]string)
	return d
}

func (d *mapToSlice) set(k, v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if rv, ok := d.data[k]; !ok {
		d.data[k] = []string{v}
	} else {
		rv = append(rv, v)
		d.data[k] = rv // MUST DO THIS
	}
}

func (d *mapToSlice) exist(k, v string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if rv, ok := d.data[k]; !ok {
		return false
	} else {
		for _, vv := range rv {
			if vv == v {
				return true
			}
		}
	}
	return false
}

func (d *mapToSlice) delete(k, v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if rv, ok := d.data[k]; ok {
		for i, r := range rv {
			if r == v {
				size := len(rv)
				copy(rv[i:], rv[i+1:])
				rv[size-1] = ""
				rv = rv[:size-1 : size-1]
				d.data[k] = rv // MUST DO THIS
				break
			}
		}
	}
}

func newMapToMap() *mapToMap {
	d := &mapToMap{}
	d.data = make(map[string]map[string]struct{})
	return d
}

func (d *mapToMap) set(k, v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if rv, ok := d.data[k]; !ok {
		td := make(map[string]struct{})
		td[v] = struct{}{}
		d.data[k] = td
	} else {
		rv[v] = struct{}{}
	}
}

func (d *mapToMap) exist(k, v string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if rv, ok := d.data[k]; !ok {
		return false
	} else {
		if _, ok := rv[v]; ok {
			return true
		}
	}
	return false
}

func (d *mapToMap) delete(k, v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if rv, ok := d.data[k]; ok {
		if _, ok := rv[v]; ok {
			delete(rv, v)
		}
	}
}
