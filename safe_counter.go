package safe_data_structures

import "sync"

type Сount interface {
	Increment()
	GetValue() int
}

type Counter struct {
	value int // значение счетчика
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
