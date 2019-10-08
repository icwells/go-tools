// Defines string based dataframe

package strarray

import (
	"github.com/icwells/go-tools/iotools"
	"strings"
	"strconv"
)

type Dataframe struct {
	Rows	[][]string
	Header	map[string]int
	Index	map[string]int
	rowlen	int
}

func NewDataFrame() *Dataframe {
	// Initializes empty struct
	var d Dataframe
	return &d
}

func DataFrameFromFile(infile string, index bool) *Dataframe {
	// Reads dataframe from text file
	d := NewDataFrame()
	tmp, d.Header = iotools.ReadFile(infile, true)
	if index {
		d.Index = make(map[string]int)
		for idx, i := range tmp {
			d.Index[i[0]] = idx
			d.Rows = append(d.Rows, i[1:])
		}
	} else {
		d.rows = tmp
	}
	d.rowlen = len(d.Rows[0])
	return d
}

func (d *Dataframe) GetHeader() []string {
	// Returns header as string slice
}

func (d *Dataframe) GetIndex(unique bool) []string {
	// Returns index as string slice
}

func (d *Dataframe) UpdateCell(idx interface, c string, v string) string {
	// Replaces given cell with v
}

func (d *Dataframe) GetCell(idx interface, c string) string {
	// Returns given cell from dataframe
}

func (d *Dataframe) GetCellInt(idx interface, c string) string {
	// Converts given cell to int
}

func (d *Dataframe) GetCellFloat(idx interface, c string) string {
	// Converts given cell to float
}

func (d *Dataframe) GetRow(idx interface) []string {
	// Returns given row from dataframe
}

func (d *Dataframe) GetColumn(c string) []string {
	// Returns given column from dataframe
	var ret []string
	for _, i := range d.Rows {
		ret = append(ret, i[d.Header[c]])
	}
	return ret
}

func (d *Dataframe) GetColumnUnique(c string) []string {
	// Returns unique values from given column
	ret := NewSet()
	for _, i := range d.Rows {
		ret.Add(i[d.Header[c]])
	}
	return ret.ToSlice()
}

