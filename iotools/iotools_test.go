// This script will perform white box tests on iotools

package iotools

import (
	"fmt"
	"testing"
)

type testcase struct {
	file     string
	delim    string
	filename string
	parent   string
	ext      string
	exists   bool
}

func newTestCases() []testcase {
	// Returns slice of test cases
	ret := []testcase{
		{"testData/comma.csv", ",", "comma", "testData", "csv", true},
		{"testData/tab.tsv", "\t", "tab", "testData", "tsv", true},
		{"testData/space.txt", " ", "space", "testData", "txt", true},
		{"test/notHere", "", "", "notHere", "", false},
	}
	return ret
}

func TestExists(t *testing.T) {
	// Tests Exists function
	cases := newTestCases()
	for _, i := range cases {
		actual := Exists(i.file)
		if actual != i.exists {
			msg := fmt.Sprintf("%s incorrectely identified by Exists", i.file)
			t.Error(msg)
		}
	}
}

func TestGetExt(t *testing.T) {
	// Tests GetExt function
	cases := newTestCases()
	for _, i := range cases {
		actual := GetExt(i.file)
		if actual != i.ext {
			msg := fmt.Sprintf("Extension for %s identified as %s", i.file, i.ext)
			t.Error(msg)
		}
	}
}

func TestGetFileName(t *testing.T) {
	// Tests GetFileName function
	cases := newTestCases()
	for _, i := range cases {
		actual := GetFileName(i.file)
		if actual != i.filename {
			msg := fmt.Sprintf("Name for %s identified as %s", i.file, i.filename)
			t.Error(msg)
		}
	}
}

func TestGetParent(t *testing.T) {
	// Tests GetParent
	cases := newTestCases()
	for _, i := range cases {
		actual := GetParent(i.file)
		if actual != i.parent {
			fmt.Println(actual, i.parent)
			msg := fmt.Sprintf("Parent directory for %s identified as %s", i.file, i.parent)
			t.Error(msg)
		}
	}
}

func TestGetDelim(t *testing.T) {
	// Tests GetDelim function
	cases := newTestCases()
	for _, i := range cases {
		first := true
		if Exists(i.file) == true {
			f := OpenFile(i.file)
			defer f.Close()
			scanner := GetScanner(f)
			for scanner.Scan() {
				if first == true {
					actual := GetDelim(string(scanner.Text()))
					if actual != i.delim {
						msg := fmt.Sprintf("Delimiter for %s identified as %s", i.file, i.delim)
						t.Error(msg)
					}
				} else {
					break
				}
			}
		}
	}
}
