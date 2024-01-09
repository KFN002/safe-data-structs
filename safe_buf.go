package safe_data_structures

import (
	"sync"
)

var (
	mu  sync.Mutex
	Buf []int
)

func Write(num int) {
	mu.Lock()
	defer mu.Unlock()
	Buf = append(Buf, num)
}

func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	if len(Buf) == 0 {
		return -1
	}
	value := Buf[0]
	Buf = Buf[1:]
	return value
}
