package gox

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
	for k := range s.items {
		sl = append(sl, k)
	}

	return sl
}

type Int64Set struct {
	items map[int64]void
}

func NewInt64Set(capacity int) *Int64Set {
	s := &Int64Set{}
	s.items = make(map[int64]void, capacity)
	return s
}

func (s *Int64Set) Add(item int64) {
	s.items[item] = void{}
}

func (s *Int64Set) Contains(item int64) bool {
	_, found := s.items[item]
	return found
}

func (s *Int64Set) Remove(item int64) {
	delete(s.items, item)
}

func (s *Int64Set) Slice() []int64 {
	sl := make([]int64, 0, len(s.items))
	for k := range s.items {
		sl = append(sl, k)
	}

	return sl
}

type StringSet struct {
	items map[string]void
}

func NewStringSet(capacity int) *StringSet {
	s := &StringSet{}
	s.items = make(map[string]void, capacity)
	return s
}

func (s *StringSet) Add(item string) {
	s.items[item] = void{}
}

func (s *StringSet) Contains(item string) bool {
	_, found := s.items[item]
	return found
}

func (s *StringSet) Remove(item string) {
	delete(s.items, item)
}

func (s *StringSet) Slice() []string {
	sl := make([]string, 0, len(s.items))
	for k := range s.items {
		sl = append(sl, k)
	}

	return sl
}
