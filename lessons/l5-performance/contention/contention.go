package contention

import (
	"sync"
	"sync/atomic"
)

type Incr struct {
	value int64
	mu    sync.Mutex
	rwmu  sync.RWMutex
}

func (i *Incr) ReadAtomic() int64 {
	return atomic.LoadInt64(&i.value)
}

func (i *Incr) IncrAtomic() {
	atomic.AddInt64(&i.value, 1)
}

func (i *Incr) ReadMutex() int64 {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.value
}

func (i *Incr) IncrMutex() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.value++
}

func (i *Incr) ReadRWMutex() int64 {
	i.rwmu.RLock()
	defer i.rwmu.RUnlock()
	return i.value
}

func (i *Incr) IncrRWMutex() {
	i.rwmu.Lock()
	defer i.rwmu.Unlock()
	i.value++
}
