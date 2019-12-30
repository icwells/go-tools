// Defines string based dataframe

package dataframe

import (
	"fmt"
	"github.com/icwells/go-tools/iotools"
	"strings"
)

// Stores data as a two-dimensional string slice with header and index as maps or named indeces.
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

// AddRow adds a string slice to dataframe and stores the index value in Index if using.
func (d *Dataframe) AddRow(row []string) error {
	var err error
	index, r := d.subsetRow(row)
	if len(r) == d.ncol {
		d.Rows = append(d.Rows, r)
		if index != "" {
			if _, ex := d.Index[index]; !ex {
				d.Index[index] = d.nrow
			} else {
				err = fmt.Errorf("Value %s already found in index.", index)
			}
		}
		d.nrow++
	} else {
		err = fmt.Errorf("Row length %d does not equal number of of columns %d.", len(r), d.ncol)
	}
	return err
}

// SetHeader converts the given row to header map.
func (d *Dataframe) SetHeader(row []string) error {
	var err error
	if d.iname != "" {
		// Determine index column number
		d.Header = iotools.GetHeader(row)
		idx, ex := d.Header[d.iname]
		if ex {
			d.col = idx
		} else {
			err = fmt.Errorf("Value %s cannot be found in header.", d.iname)
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
		err = fmt.Errorf("%v is not a valid header index. Must be string or integer.", i)
	}
	return err
}

// NewDataFrame returns an empty dataframe struct.
func NewDataFrame(column interface{}) (*Dataframe, error) {
	d := new(Dataframe)
	err := d.setIndexColumn(column)
	d.Header = make(map[string]int)
	d.Index = make(map[string]int)
	return d, err
}

// DataFrameFromFile creates a dataframe and loads in data from the given input file. The first row is assumed to be the header.
func FromFile(infile string, column interface{}) (*Dataframe, error) {
	var tmp [][]string
	d, err := NewDataFrame(column)
	tmp, d.Header = iotools.ReadFile(infile, true)
	d.ncol = len(d.Header)
	if d.col >= 0 {
		// Remove index column from header
		err = d.SetHeader(d.GetHeader())
	}
	for _, i := range tmp {
		err = d.AddRow(i)
		if err != nil {
			break
		}
	}
	return d, err
}

// ToSlice returns dataframe as a two-dimensional string slice (inserts index values if needed but does not include header).
func (d *Dataframe) ToSlice() [][]string {
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

// ToCSV writes datframe to csv with header and index inderted.
func (d *Dataframe) ToCSV(outfile string) {
	h := strings.Join(d.GetHeader(), ",")
	if d.col >= 0 {
		// Prepend index column name
		h = d.iname + "," + h
	}
	iotools.WriteToCSV(outfile, h, d.ToSlice())
}
