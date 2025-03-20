package rust_gc

import (
	"sync"
)

type Gc struct {
	mu      sync.Mutex
	objects map[*Object]int
}

type Object struct {
	value any
}

func NewGc() *Gc {
	return &Gc{
		objects: make(map[*Object]int),
	}
}

func (gc *Gc) Allocate(value any) *Object {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	obj := &Object{value: value}
	gc.objects[obj] = 1
	return obj
}

func (gc *Gc) AddRef(obj *Object) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	if _, exists := gc.objects[obj]; exists {
		gc.objects[obj]++
	}
}

func (gc *Gc) Release(obj *Object) {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	if count, exists := gc.objects[obj]; exists {
		if count > 1 {
			gc.objects[obj]--
		} else {
			delete(gc.objects, obj)
		}
	}
}

func (gc *Gc) Collect() {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	for obj, count := range gc.objects {
		if count == 0 {
			delete(gc.objects, obj)
		}
	}
}
