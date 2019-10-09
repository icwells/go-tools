// Tests dataframe package

package dataframe

import (
	"strings"
	"testing"
)

func getTestSlice() [][]string {
	// Returns slice for testing data frame
	return [][]string{
		{"Sex", "Age", "ID", "Species", "Name"},
		{"male", "24", "1", "Heloderma suspectum", "Gila monster"},
		{"NA", "31.0", "2", "Canis lupus", "wolf"},
		{"female", "12.1", "3", "Canis latrans", "coyote"},
		{"male", "2", "4", "Canis aureus", "jackal"},
	}
}

func evaluateDF(t *testing.T, df *Dataframe, rows [][]string, index int) {
	// Compares dataframe to rows
	h := len(rows[0])
	if index >= 0 {
		h--
	}
	c, r := df.Dimensions()
	if r != len(rows) - 1 {
		t.Errorf("Dimensions returned %d rows instead of %d", r, len(rows) - 1)
	}
	if c != h {
		t.Errorf("Dimensions returned %d columns instead of %d", c, h)
	}
	head := strings.Join(df.GetHeader(), " ")
	idx, ehead := df.subsetRow(rows[0])
	eh := strings.Join(ehead, " ")
	if head != eh {
		t.Errorf("Actual header %s does not equal expected: %s", head, eh)
	}
	if idx != df.iname {
		t.Errorf("Actual index name %s does not equal expected: %s", df.iname, idx)
	}	
}

func TestDataFrame(t *testing.T) {
	rows := getTestSlice()
	for _, i := range []int{-1, 2} {
		df := NewDataFrame(i)
		df.SetHeader(rows[0])
		for _, i := range rows[1:] {
			err := df.AddRow(i)
			if err != nil {
				t.Errorf("Error setting dataframe row: %v", err)
			}
		}
		evaluateDF(t, df, rows, i)
	}
}
