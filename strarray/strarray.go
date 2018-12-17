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

func NewSet() Set {
	// Initializes new set
	var s Set
	s.set = make(map[string]byte)
	return s
}

func (s *Set) Length() int {
	// Returns length of set
	return len(s.set)
}

func (s *Set) Add(v string) {
	// Add new value to set
	s.set[v] = '0'
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
