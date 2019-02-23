package gox

import "sync"

type void struct{}

type Set struct {
	items map[interface{}]void
}

func NewSet(capacity int) *Set {
	s := &Set{}
	s.items = make(map[interface{}]void, capacity)
	return s
}

func (s *Set) Add(item interface{}) {
	s.items[item] = void{}
}

func (s *Set) Contains(item interface{}) bool {
	_, found := s.items[item]
	return found
}

func (s *Set) Remove(item interface{}) {
	delete(s.items, item)
}

func (s *Set) Slice() []interface{} {
	sl := make([]interface{}, 0, len(s.items))
	for k, _ := range s.items {
		sl = append(sl, k)
	}

	return sl
}

type ConcurrentSet struct {
	items map[interface{}]void
	mu    sync.RWMutex
}

func NewConcurrentSet(capacity int) *ConcurrentSet {
	s := &ConcurrentSet{}
	s.items = make(map[interface{}]void, capacity)
	return s
}

func (s *ConcurrentSet) Add(item interface{}) {
	s.mu.Lock()
	s.items[item] = void{}
	s.mu.Unlock()
}

func (s *ConcurrentSet) Contains(item interface{}) bool {
	s.mu.RLock()
	_, found := s.items[item]
	s.mu.RUnlock()
	return found
}

func (s *ConcurrentSet) Remove(item interface{}) {
	s.mu.Lock()
	delete(s.items, item)
	s.mu.Unlock()
}

func (s *ConcurrentSet) Slice() []interface{} {
	s.mu.RLock()
	sl := make([]interface{}, 0, len(s.items))
	for k, _ := range s.items {
		sl = append(sl, k)
	}
	s.mu.RUnlock()
	return sl
}
