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
	if len(r) == d.ncol {
		d.Rows = append(d.Rows, r)
		if index != "" {
			if _, ex := d.Index[index]; !ex {
				d.Index[index] = d.nrow
			} else {
				err = errors.New(fmt.Sprintf("Value %s already found in index.", index))
			}
		}
		d.nrow++
	} else {
		err = errors.New(fmt.Sprintf("Row length %d does not equal number of of columns %d.", len(r), d.ncol))
	}
	return err
}

func (d *Dataframe) SetHeader(row []string) error {
	// Converts given row to header
	var err error
	if d.iname != "" {
		// Determine index column number
		d.Header = iotools.GetHeader(row)
		idx, ex := d.Header[d.iname]
		if ex {
			d.col = idx
		} else {
			err = errors.New(fmt.Sprintf("Value %s cannot be found in header.", d.iname))
		}
	}
	index, r := d.subsetRow(row)
	if index != "" && d.iname == "" {
		// Store index column name
		d.iname = index
	}
	d.Header = iotools.GetHeader(r)
	d.ncol = len(d.Header)
	return err
}

func (d *Dataframe) setIndexColumn(column interface{}) error {
	// Assigns column tp d.col/d.iname
	var err error
	switch i := column.(type) {
	case string:
		d.iname = string(i)
	case int:
		d.col = int(i)
	default:
		err = errors.New(fmt.Sprintf("%v is not a valid header index. Must be string or integer.", i))
	}
	return err
}

func NewDataFrame(column interface{}) (*Dataframe, error) {
	// Initializes empty struct
	d := new(Dataframe)
	err := d.setIndexColumn(column)
	d.Header = make(map[string]int)
	d.Index = make(map[string]int)
	return d, err
}

func DataFrameFromFile(infile string, column interface{}) (*Dataframe, error) {
	// Reads dataframe from text file
	var tmp [][]string
	d, err := NewDataFrame(column)
	tmp, d.Header = iotools.ReadFile(infile, true)
	if d.col >= 0 {
		// Remove index column from header
		err = d.SetHeader(d.GetHeader())
	}
	for _, i := range tmp {
		d.AddRow(i)
	}
	return d, err
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
