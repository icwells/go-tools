// Defines string based dataframe

package dataframe

import (
	"errors"
	"fmt"
	"github.com/icwells/go-tools/iotools"
	"strings"
)

type Dataframe struct {
	Rows   [][]string
	Header map[string]int
	Index  map[string]int
	iname  string
	col    int
	ncol   int
	nrow   int
}

func (d *Dataframe) subsetRow(row []string) (string, []string) {
	// Returns row with index column seperate
	var index string
	var ret []string
	if d.col >= 0 {
		for idx, i := range row {
			i = strings.TrimSpace(i)
			if idx == d.col {
				index = i
			} else {
				ret = append(ret, i)
			}
		}
	} else {
		ret = row
	}
	return index, ret
}

func (d *Dataframe) AddRow(row []string) error {
	// Adds row to dataframe
	var err error
	index, r := d.subsetRow(row)
	d.Rows = append(d.Rows, r)
	if index != "" {
		if _, ex := d.Index[index]; !ex {
			d.Index[index] = d.nrow
		} else {
			err = errors.New(fmt.Sprintf("Value %s already found in index.", index))
		}
	}
	d.nrow++
	return err
}

func (d *Dataframe) SetHeader(row []string) {
	// Converts given row to header
	index, r := d.subsetRow(row)
	if index != "" {
		d.iname = index
	}
	d.Header = iotools.GetHeader(r)
	d.ncol = len(d.Header)
}

func NewDataFrame(column int) *Dataframe {
	// Initializes empty struct
	d := new(Dataframe)
	d.col = column
	d.Header = make(map[string]int)
	d.Index = make(map[string]int)
	return d
}

func DataFrameFromFile(infile string, column int) *Dataframe {
	// Reads dataframe from text file
	var tmp [][]string
	d := NewDataFrame(column)
	tmp, d.Header = iotools.ReadFile(infile, true)
	if d.col >= 0 {
		// Remove index column from header
		d.SetHeader(d.GetHeader())
	}
	for _, i := range tmp {
		d.AddRow(i)
	}
	return d
}

func (d *Dataframe) ToSlice() [][]string {
	// Returns dataframe as slice of string slices (inserts index values if needed)
	var ret [][]string
	if d.col >= 0 {
		// Prepend index to rows
		index := d.GetIndex()
		for idx, i := range d.Rows {
			row := append([]string{index[idx]}, i...)
			ret = append(ret, row)
		}
	} else {
		ret = d.Rows
	}
	return ret
}

func (d *Dataframe) ToCSV(outfile string) {
	// Writes rows to csv
	h := strings.Join(d.GetHeader(), ",")
	if d.col >= 0 {
		// Prepend index column name
		h = d.iname + "," + h
	}
	iotools.WriteToCSV(outfile, h, d.ToSlice())
}
