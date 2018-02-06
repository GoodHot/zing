package util

import "sync"

// Set 集合
type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

// NewSet 创建
func NewSet() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

// Add 添加元素
func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// Remove 删除元素
func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// Has 是否存在
func (s *Set) Has(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Len 长度
func (s *Set) Len() int {
	return len(s.m)
}

// Clear 清空
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

// IsEmpty 是否为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// Each 遍历
func (s *Set) Each(f func(interface{})) {
	for k := range s.m {
		f(k)
	}
}
