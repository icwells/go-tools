// Commonly used functions for handling string-based arrays

package strarray

import (
	"strings"
)

// TitleCase manually converts term to title case (strings.Title is buggy).
func TitleCase(t string) string {
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

// InSliceStr returns true if s is in l.
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

// InSliceSli returns true if s is in column c in l.
func InSliceSli(l [][]string, s string, c int) bool {
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

// SliceIndex returns the first index of v in l.
func SliceIndex(l []string, v string) int {
	ret := -1
	for idx, i := range l {
		if i == v {
			ret = idx
			break
		}
	}
	return ret
}

// SliceCount returns the number of occurances of v in s.
func SliceCount(s []string, v string) int {
	ret := 0
	for _, i := range s {
		if i == v {
			ret++
		}
	}
	return ret
}

// DeleteSliceIndex deletes item at idx while preventing index errors.
func DeleteSliceIndex(s []string, idx int) []string {
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

// DeleteSliceValue deletes all occurances of v from s.
func DeleteSliceValue(s []string, v string) []string {
	var ret []string
	for _, i := range s {
		if i != v {
			ret = append(ret, i)
		}
	}
	return ret
}
