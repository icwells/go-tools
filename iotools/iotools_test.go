// This script will perform white box tests on iotools

package iotools

import (
	"strings"
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
			t.Errorf("%s incorrectely identified by Exists", i.file)
		}
	}
}

func TestGetExt(t *testing.T) {
	// Tests GetExt function
	cases := newTestCases()
	for _, i := range cases {
		actual := GetExt(i.file)
		if actual != i.ext {
			t.Errorf("Extension for %s identified as %s", i.file, actual)
		}
	}
}

func TestGetFileName(t *testing.T) {
	// Tests GetFileName function
	cases := newTestCases()
	for _, i := range cases {
		actual := GetFileName(i.file)
		if actual != i.filename {
			t.Errorf("Name for %s identified as %s", i.file, actual)
		}
	}
}

func TestGetParent(t *testing.T) {
	// Tests GetParent
	cases := newTestCases()
	for _, i := range cases {
		actual := GetParent(i.file)
		if actual != i.parent {
			t.Errorf("Parent directory for %s identified as %s", i.file, actual)
		}
	}
}

func TestFindDelim(t *testing.T) {
	// Tests FindDelim function
	cases := newTestCases()
	for _, i := range cases {
		if i.exists {
			act, _ := FindDelim(i.file)
			if act != i.delim {
				t.Errorf("Delimiter for %s identified as %s", i.file, act)
				break
			}
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
					actual, _ := GetDelim(string(scanner.Text()))
					if actual != i.delim {
						t.Errorf("Delimiter for %s identified as %s", i.file, actual)
					}
				} else {
					break
				}
			}
		}
	}
}

func testSlice() [][]string {
	return [][]string{
		{"Sex", "Age", "Castrated", "ID", "Species", "Date", "Comments", "MassPresent", "Necropsy", "Metastasis", "TumorType", "Location", "Primary", "Malignant", "Service", "Account", "Submitter"},
		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17"},
		{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18"},
	}
}

func TestYieldFile(t *testing.T) {
	var idx int
	exp := testSlice()
	act, _ := YieldFile("testData/test.csv", false)
	for i := range act {
		a := strings.Join(i, ",")
		e := strings.Join(exp[idx], ",")
		if a != e {
			t.Errorf("Actual line %s does not equal expected: %s", a, e)
		}
		idx++
	}
}

func TestReadFile(t *testing.T) {
	exp := testSlice()
	act, _ := ReadFile("testData/test.csv", false)
	for idx, i := range act {
		a := strings.Join(i, ",")
		e := strings.Join(exp[idx], ",")
		if a != e {
			t.Errorf("Actual line %s does not equal expected: %s", a, e)
		}
	}
}
