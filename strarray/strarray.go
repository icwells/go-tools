// Commonly used functions for handling string-based arrays

package strarray

import (
	"sort"
)

func InSliceStr(l []string, s string) bool {
	// Returns true if s is in l
	in := false
	for _, i := range l {
		if s == i {
			in = true
			break
		}
	}
	return in
}

func InSliceSli(l [][]string, s string, c int) bool {
	// Returns true if s is in column c in l
	in := false
	for _, i := range l {
		if c < len(i) {
			if s == i[c] {
				in = true
				break
			}
		}
	}
	return in
}

//----------------------------------------------------------------------------

type Set struct {
	set	map[string]byte
}

func NewSet() {
	// Initializes new set
	return make(map[string]byte)
}

func (s *Set) Len() int {
	// Returns length of set
	return len(s.set)
}

func (s *Set) Add(v string) {
	// Add new value to set
	s[v] = '0'
}

func (s *Set) InSet(v string) {
	// Returns true if v is in s
	_, ex := s.set[v]
	return ex
}

func (s *Set) ToSlice() {
	// Returns sorted slice of set
	var ret []string
	for k := range s.set {
		ret = append(ret, k)
	}
	sort.Sort(ret)
	return ret
}

func (s *Set) Pop(v string) {
	// Removes v from set
	if s.InSet(v) == true {
		delete(s.set, v)
	}
}
