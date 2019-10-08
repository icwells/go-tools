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
	d.Header = make(map[string]int)
	d.Index = make(map[string]int)
	return &d
}

func DataFrameFromFile(infile string, index bool) *Dataframe {
	// Reads dataframe from text file
	d := NewDataFrame()
	tmp, d.Header = iotools.ReadFile(infile, true)
	if index {
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

func (d *Dataframe) ToCSV(outfile string) {
	// Writes rows to csv
	var tmp [][]string
	if len(d.Index) > 0 {
		index := d.GetIndex()
		for idx, i := range d.Rows {
			row = append([]string{index[idx]}, i...)
			tmp = append(tmp, row)
		}	
	} else {
		tmp = d.Rows
	}
	h := strings.Join(d.GetHeader(), ",")
	iotools.WriteToCSV(outfile, h, tmp)
}

func (d *Dataframe) getIndeces(idx interface{}, col) (int, int) {
	// Returns integer index for rows
	var r int
	if idx.(type) == string && len(d.Index) > 0 {
		r = d.Index[string(idx)]
	} else {
		r = int(idx)
	}
	c, ex := d.Header[col]
	if ex == false {
		// Print error?
	}
	return r, c
}

func (d *Dataframe) UpdateCell(idx interface{}, col string, v string) string {
	// Replaces given cell with v
	r, c := d.getIndeces(idx, col)
	d.Rows[r][c] = v
}

func (d *Dataframe) GetCell(idx interface{}, col string) string {
	// Returns given cell from dataframe
	r, c := d.getIndeces(idx, col)
	return d.Rows[r][c]
}

func (d *Dataframe) GetCellInt(idx interface{}, col string) string {
	// Converts given cell to int
	val := GetCell(idx, col)
	
}

func (d *Dataframe) GetCellFloat(idx interface{}, col string) string {
	// Converts given cell to float
	val := GetCell(idx, col)
}

func (d *Dataframe) GetRow(idx interface{}) []string {
	// Returns given row from dataframe
}

func (d *Dataframe) GetColumn(col string) []string {
	// Returns given column from dataframe
	var ret []string
	for _, i := range d.Rows {
		ret = append(ret, i[d.Header[c]])
	}
	return ret
}

func (d *Dataframe) GetColumnUnique(col string) []string {
	// Returns unique values from given column
	ret := NewSet()
	for _, i := range d.Rows {
		ret.Add(i[d.Header[c]])
	}
	return ret.ToSlice()
}

