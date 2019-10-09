// Defines python-like string set

package strarray

import (
	"sort"
)

type Set struct {
	set map[string]byte
}

func NewSet() Set {
	// Initializes new set
	var s Set
	s.set = make(map[string]byte)
	return s
}

func ToSet(v []string) Set {
	// Converts string slice to set
	ret := NewSet()
	for _, i := range v {
		ret.Add(i)
	}
	return ret
}

func (s *Set) Length() int {
	// Returns length of set
	return len(s.set)
}

func (s *Set) Add(v string) {
	// Add new value to set
	s.set[v] = '0'
}

func (s *Set) Extend(v []string) {
	// Adds all elements of slice to set
	for _, i := range v {
		s.Add(i)
	}
}

func (s *Set) InSet(v string) bool {
	// Returns true if v is in s
	_, ex := s.set[v]
	return ex
}

func (s *Set) ToSlice() []string {
	// Returns sorted slice of set
	var ret []string
	for k := range s.set {
		ret = append(ret, k)
	}
	sort.Strings(ret)
	return ret
}

func (s *Set) Pop(v string) {
	// Removes v from set
	if s.InSet(v) == true {
		delete(s.set, v)
	}
}
