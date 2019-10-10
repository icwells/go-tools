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

func evaluateGetRow(t *testing.T, df *Dataframe, rows [][]string) {
	// Tests getRows
	for idx, i := range rows[1:] {
		_, e := df.subsetRow(i)
		exp := strings.Join(e, " ")
		a, err := df.GetRow(idx)
		if err != nil {
			t.Errorf("Error selecting row %d: %v", idx, err)
		} else {
			act := strings.Join(a, " ")
			if act != exp {
				t.Errorf("Actual row %s does not equal expected: %s", act, exp)
			}
		}
	}
}

func evaluateGetColumn(t *testing.T, df *Dataframe, rows [][]string) {
	// Tests getColumn functions
	var e []string
	for _, i := range rows[1:] {
		e = append(e, i[1])
	}
	exp := strings.Join(e, " ")
	a, err := df.GetColumn("Age")
	if err != nil {
		t.Errorf("Error selecting column Age: %v", err)
	} else {
		act := strings.Join(a, " ")
		if act != exp {
			t.Errorf("Actual column %s does not equal expected: %s", act, exp)
		}
	}
}

func evaluateGetCell(t *testing.T, df *Dataframe, exp string, index int) {
	// Tests get cellfor each data type
	var act string
	var err error
	var r interface{}
	c := "Age"
	if index >= 0 {
		r = "3"
	} else {
		r = 2
	}
	act, err = df.GetCell(r, c)
	if err != nil {
		t.Errorf("Error selecting cell at %v, %s: %v", r, c, err)
	}
	if act != exp {
		t.Errorf("Actual cell value %s does not equal expected: %s", act, exp)
	}
	ai, err := df.GetCellInt(r, c)
	if err != nil {
		t.Errorf("Error selecting cell at %v, %s: %v", r, c, err)
	}
	if ai != 12 {
		t.Errorf("Actual cell value %d does not equal expected: 12", ai)
	}
	af, err := df.GetCellFloat(r, c)
	if err != nil {
		t.Errorf("Error selecting cell at %v, %s: %v", r, c, err)
	}
	if af != 12.1 {
		t.Errorf("Actual cell value %f does not equal expected: 12.1", af)
	}
}

func evaluateHeader(t *testing.T, df *Dataframe, rows [][]string) {
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

func evaluateIndex(t *testing.T, df *Dataframe, rows [][]string) {
	// Compares indeces
	var e []string
	index := strings.Join(df.GetIndex(), " ")
	for _, i := range rows[1:] {
		e = append(e, i[2])
	}
	exp := strings.Join(e, " ")
	if index != exp {
		t.Errorf("Actual index %s does not equal expected: %s", index, exp)
	}
}

func evaluateDF(t *testing.T, df *Dataframe, rows [][]string, index int) {
	// Compares dataframe to rows
	h := len(rows[0])
	if index >= 0 {
		h--
		evaluateIndex(t, df, rows)
	}
	c, r := df.Dimensions()
	if r != len(rows)-1 {
		t.Errorf("Dimensions returned %d rows instead of %d", r, len(rows)-1)
	}
	if c != h {
		t.Errorf("Dimensions returned %d columns instead of %d", c, h)
	}
	evaluateHeader(t, df, rows)
	evaluateGetCell(t, df, "12.1", index)
	evaluateGetRow(t, df, rows)
	evaluateGetColumn(t, df, rows)
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
