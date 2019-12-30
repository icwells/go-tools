// Getter methods for dataframe

package dataframe

import (
	"errors"
	"fmt"
	"github.com/icwells/simpleset"
	"strconv"
	"strings"
)

func (d *Dataframe) Dimensions() (int, int) {
	// Returns row and column lengths
	return d.ncol, d.nrow
}

func (d *Dataframe) Length() int {
	// Returns number of rows
	return d.nrow
}

func (d *Dataframe) GetHeader() []string {
	// Returns header as string slice
	ret := make([]string, len(d.Header))
	for k, v := range d.Header {
		ret[v] = k
	}
	return ret
}

func (d *Dataframe) GetIndex() []string {
	// Returns index as string slice
	ret := make([]string, len(d.Index))
	for k, v := range d.Index {
		ret[v] = k
	}
	return ret
}

func (d *Dataframe) getIndex(m map[string]int, l int, name string, idx interface{}) (int, error) {
	// Returns index for column or row
	var ex bool
	var err error
	var ret int
	switch i := idx.(type) {
	case string:
		if len(m) > 0 {
			ret, ex = m[string(i)]
			if !ex {
				err = errors.New(fmt.Sprintf("%s name %v cannot be found.", name, i))
			}
		} else {
			err = errors.New(fmt.Sprintf("String indeces cannot be used if dataframe %s is not set.", strings.ToLower(name)))
		}
	case int:
		if i < l {
			ret = int(i)
		} else {
			err = errors.New(fmt.Sprintf("Integer index %v exceeds %s length %d.", i, strings.ToLower(name), l))
		}
	default:
		err = errors.New(fmt.Sprintf("%v is not a valid index. Must be string or integer.", i))
	}
	return ret, err
}

func (d *Dataframe) getIndeces(idx interface{}, col interface{}) (int, int, error) {
	// Returns integer index for rows
	var c int
	r, err := d.getIndex(d.Index, d.nrow, "Index", idx)
	if err == nil {
		c, err = d.getIndex(d.Header, d.ncol, "Header", col)
	}
	return r, c, err
}

func (d *Dataframe) UpdateCell(idx interface{}, col interface{}, v string) error {
	// Replaces given cell with v
	r, c, err := d.getIndeces(idx, col)
	if err == nil {
		d.Rows[r][c] = v
	}
	return err
}

func (d *Dataframe) GetCell(idx interface{}, col interface{}) (string, error) {
	// Returns given cell from dataframe
	var ret string
	r, c, err := d.getIndeces(idx, col)
	if err == nil {
		ret = d.Rows[r][c]
	}
	return ret, err
}

func (d *Dataframe) GetCellInt(idx interface{}, col interface{}) (int, error) {
	// Converts given cell to int
	var ret int
	val, err := d.GetCell(idx, col)
	if err == nil {
		// Remove any decimal values
		val = strings.Split(val, ".")[0]
		ret, err = strconv.Atoi(val)
	}
	return ret, err
}

func (d *Dataframe) GetCellFloat(idx interface{}, col interface{}) (float64, error) {
	// Converts given cell to float64
	var ret float64
	val, err := d.GetCell(idx, col)
	if err == nil {
		ret, err = strconv.ParseFloat(val, 64)
	}
	return ret, err
}

func (d *Dataframe) GetRow(idx interface{}) ([]string, error) {
	// Returns given row from dataframe
	var ret []string
	r, err := d.getIndex(d.Index, d.nrow, "Index", idx)
	if err == nil {
		ret = d.Rows[r]
	}
	return ret, err
}

func (d *Dataframe) SliceRow(idx interface{}, start interface{}, end interface{}) ([]string, error) {
	// Reutrns row[start:end]
	var ret []string
	r, s, err := d.getIndeces(idx, start)
	if err == nil {
		var e int
		e, err = d.getIndex(d.Header, d.ncol, "Header", end)
		if e < 0 {
			e = d.ncol
			err = nil
		}
		if err == nil {
			if e > s {
				if e-s == 1 {
					ret = append(ret, d.Rows[r][s])
				} else {
					ret = d.Rows[r][s:e]
				}
			} else {
				err = errors.New(fmt.Sprintf("Starting index %d is greater than %d.", s, e))
			}
		}
	}
	return ret, err
}

func (d *Dataframe) GetColumn(col interface{}) ([]string, error) {
	// Returns given column from dataframe
	var ret []string
	c, err := d.getIndex(d.Header, d.ncol, "Header", col)
	if err == nil {
		for _, i := range d.Rows {
			ret = append(ret, i[c])
		}
	}
	return ret, err
}

func (d *Dataframe) GetColumnUnique(col interface{}) ([]string, error) {
	// Returns unique values from given column
	ret := simpleset.NewStringSet()
	c, err := d.getIndex(d.Header, d.ncol, "Header", col)
	if err == nil {
		for _, i := range d.Rows {
			ret.Add(i[c])
		}
	}
	return ret.ToStringSlice(), err
}
