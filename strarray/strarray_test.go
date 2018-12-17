// This script will perform white box tests on strarray

package strarray

import (
	"testing"
)

type testcase struct {
	slice 		[]string
	target		string
	column		int
	expected	bool
}

func newStringSliceCases() []testcase {
	// Returns slice of cases
	ret := []testcase{
		{[]string{"a", "b", "c"}, "b", 0, true},
		{[]string{"a", "b", "c"}, "B", 0, false},
		{[]string{"x", "y", "z", "seven"}, "seven", 0, true},
		{[]string{}, "b", 0, false},
	}
	return ret
}

func TestInSliceStr(t *testing.T) {
	// Tests InSliceStr function
	cases := newStringSliceCases()
	for _, i := range cases {
		actual := InSliceStr(i.slice, i.target)
		if actual != i.expected {
			t.Errorf("%s: %s", i.target, i.slice)
		}
	}
}

func newSliceSliceCases() []testcase {
	// Returns test cases
	ret := []testcase{
		{[]string{}, "b", 1, true},
		{[]string{}, "B", 0, false},
		{[]string{}, "six", 2, true},
		{[]string{}, "b", 5, false},
	}
	return ret
}

func TestInSliceSlice(t *testing.T) {
	// Tests InSliceSli function
	slice := [][]string{
		{"a", "b", "C"},
		{"x", "y", "z"},
		{"monitor", "trackball", "six"},
	}
	cases := newSliceSliceCases()
	for _, i := range cases {
		actual := InSliceSli(slice, i.target, i.column)
		if actual != i.expected {
			t.Errorf("%s: %s", i.target, i.slice)
		}
	}
}

func TestSliceIndex(t *testing.T) {
	// Tests slice index function
	cases := []string{"a", "b", "c", "d", "e"}
	for idx, i := range cases {
		a := SliceIndex(cases, i)
		if a != idx {
			t.Errorf("Actual index %d does not equal expected: %d", a, idx)
		}
	}
	a := SliceIndex(cases, "f")
	if a != -1 {
		t.Errorf("Actual index %d does not equal -1", a)
	}
}

func evaluateLength(t *testing.T, a, e int) {
	// Compares actual length to expected
	if a != e {
		t.Errorf("Actual length %d does nto equal expected: %d", a, e)
	}
}

func evaluateBool(t *testing.T, a, e bool) {
	// Compares results of inset
	if a != e {
		t.Errorf("Actual InSet value %v does nto equal expected: %v", a, e)
	}
}

func TestSet(t *testing.T) {
	// Tests set attributes
	cases := []string{"a", "b", "c", "d", "e"}
	s := NewSet()
	evaluateLength(t, s.Length(), 0)
	for idx, i := range cases {
		s.Add(i)
		evaluateLength(t, s.Length(), idx+1)
		evaluateBool(t, s.InSet(i), true)
	}
	l := s.Length()
	for _, i := range cases {
		s.Pop(i)
		l--
		evaluateLength(t, s.Length(), l)
		evaluateBool(t, s.InSet(i), false)
	}
}
