package safe_data_structures

import "sync"

type Queue interface {
	Enqueue(element interface{})
	Dequeue() interface{}
}

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

func (cq *ConcurrentQueue) Enqueue(element interface{}) {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	cq.queue = append(cq.queue, element)
}

func (cq *ConcurrentQueue) Dequeue() interface{} {
	cq.mutex.Lock()
	defer cq.mutex.Unlock()
	if len(cq.queue) == 0 {
		return nil
	}
	element := cq.queue[0]
	cq.queue = cq.queue[1:]
	return element
}
