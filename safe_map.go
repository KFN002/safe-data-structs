package safe_data_structures

import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.m[key]
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[key] = value
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m:   make(map[string]interface{}),
		mux: sync.Mutex{},
	}
}
