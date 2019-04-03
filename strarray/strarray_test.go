// This script will perform white box tests on strarray

package strarray

import (
	"testing"
)

func TestTitleCase(t *testing.T) {
	str := []struct {
		input, expected string
	}{
		{"SEBA'S STRIPED  FINGERFISH", "Seba's Striped Fingerfish"},
		{"Sharp shinned Hawk", "Sharp Shinned Hawk"},
		{"PIPING` x GUAN ", "Piping` Guan"},
	}
	for _, i := range str {
		a := TitleCase(i.input)
		if a != i.expected {
			t.Errorf("Actual term %s does not equal expected: %s", a, i.expected)
		}
	}
}

type testcase struct {
	slice    []string
	target   string
	column   int
	expected bool
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

func TestSliceCount(t *testing.T) {
	// Tests SliceCount
	s := []string{"a", "b", "e", "e", "c", "c", "d", "e"}
	cases := []struct {
		v string
		i int
	}{
		{"a", 1},
		{"b", 1},
		{"e", 3},
		{"c", 2},
		{"d", 1},
	}
	for _, i := range cases {
		n := SliceCount(s, i.v)
		if n != i.i {
			t.Errorf("Incorrect count %d returned for %s.", n, i.v)
		}
	}
}

func TestDeleteSliceIndex(t *testing.T) {
	//Tests DeleteSliceIndex function
	cases := []string{"a", "b", "c", "d", "e"}
	for idx, i := range cases {
		cp := make([]string, len(cases))
		copy(cp, cases)
		cp = DeleteSliceIndex(cp, idx)
		if idx < len(cp) && cp[idx] == i {
			t.Errorf("Value %s at index %d was not removed from slice.", i, idx)
		}
	}
}

func TestDeleteSliceValue(t *testing.T) {
	//Tests DeleteSliceValue function
	cases := []string{"a", "b", "e", "e", "c", "c", "d", "e"}
	for _, i := range cases {
		cp := make([]string, len(cases))
		copy(cp, cases)
		cp = DeleteSliceValue(cp, i)
		if SliceCount(cp, i) != 0 {
			t.Errorf("All instances of %s were not removed from slice.", i)
		}
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
