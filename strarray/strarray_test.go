// This script will perform white box tests on strarray

package strarray

import (
	"fmt"
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
			msg := fmt.Sprintf("%s: %s", i.target, i.slice)
			t.Error(msg)
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
			fmt.Println(actual, i.expected)
			msg := fmt.Sprintf("%s: %s", i.target, i.slice)
			t.Error(msg)
		}
	}
}
