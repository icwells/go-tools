// Commonly used functions for handling string-based arrays

package strarray

import (
	"sort"
	"strings"
)

func TitleCase(t string) string {
	// Manually converts term to title case (strings.Title is buggy)
	var query []string
	s := strings.Split(t, " ")
	for _, i := range s {
		if len(i) > 1 {
			// Skip stray characters
			query = append(query, strings.ToUpper(string(i[0]))+strings.ToLower(i[1:]))
		}
	}
	return strings.Join(query, " ")
}

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

func SliceIndex(l []string, v string) int {
	// Returns first index of v in l
	ret := -1
	for idx, i := range l {
		if i == v {
			ret = idx
			break
		}
	}
	return ret
}

func SliceCount(s []string, v string) int {
	// Returns number of occurances of v in s
	ret := 0
	for _, i := range s {
		if i == v {
			ret++
		}
	}
	return ret
}

func DeleteSliceIndex(s []string, idx int) []string {
	// Deletes item at idx while preventing index errors
	if len(s) > 0 {
		if idx == 0 {
			s = s[idx+1:]
		} else if idx == len(s)-1 {
			s = s[:idx]
		} else if idx < len(s) {
			s = append(s[:idx], s[idx+1:]...)
		}
	}
	return s
}

func DeleteSliceValue(s []string, v string) []string {
	// Deletes all occurances of v from s
	var ret []string
	for _, i := range s {
		if i != v {
			ret = append(ret, i)
		}
	}
	return ret
}

//----------------------------------------------------------------------------

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
