// Getter methods for dataframe

package dataframe

import (
	"errors"
	"fmt"
	"github.com/icwells/simpleset"
	"strconv"
	"strings"
)

// Dimensions returns row and column lengths (respectively).
func (d *Dataframe) Dimensions() (int, int) {
	return d.ncol, d.nrow
}

// Length returns the number of rows.
func (d *Dataframe) Length() int {
	return d.nrow
}

// GetHeader returns the header as string slice.
func (d *Dataframe) GetHeader() []string {
	ret := make([]string, len(d.Header))
	for k, v := range d.Header {
		ret[v] = k
	}
	return ret
}

// GetIndex returns index as string slice.
func (d *Dataframe) GetIndex() []string {
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

// UpdateCell replaces the value in a given cell with v.
func (d *Dataframe) UpdateCell(idx interface{}, col interface{}, v string) error {
	r, c, err := d.getIndeces(idx, col)
	if err == nil {
		d.Rows[r][c] = v
	}
	return err
}

// GetCell returns the value of the given cell from dataframe.
func (d *Dataframe) GetCell(idx interface{}, col interface{}) (string, error) {
	var ret string
	r, c, err := d.getIndeces(idx, col)
	if err == nil {
		ret = d.Rows[r][c]
	}
	return ret, err
}

// GetCellInt returns the value of the given cell as int. Returns an error if it cannot be converted.
func (d *Dataframe) GetCellInt(idx interface{}, col interface{}) (int, error) {
	var ret int
	val, err := d.GetCell(idx, col)
	if err == nil {
		// Remove any decimal values
		val = strings.Split(val, ".")[0]
		ret, err = strconv.Atoi(val)
	}
	return ret, err
}

// GetCellFloat returns the value of the given cell as float64. Returns an error if it cannot be converted.
func (d *Dataframe) GetCellFloat(idx interface{}, col interface{}) (float64, error) {
	var ret float64
	val, err := d.GetCell(idx, col)
	if err == nil {
		ret, err = strconv.ParseFloat(val, 64)
	}
	return ret, err
}

// GetRow returns given row from dataframe. Does not insert index value.
func (d *Dataframe) GetRow(idx interface{}) ([]string, error) {
	var ret []string
	r, err := d.getIndex(d.Index, d.nrow, "Index", idx)
	if err == nil {
		ret = d.Rows[r]
	}
	return ret, err
}

// SliceRow returns a slice of the given row between start and end (not including the value at end).
func (d *Dataframe) SliceRow(idx interface{}, start interface{}, end interface{}) ([]string, error) {
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

// GetColumn returns the given column from dataframe.
func (d *Dataframe) GetColumn(col interface{}) ([]string, error) {
	var ret []string
	c, err := d.getIndex(d.Header, d.ncol, "Header", col)
	if err == nil {
		for _, i := range d.Rows {
			ret = append(ret, i[c])
		}
	}
	return ret, err
}

// GetColumnUnique returns all unique values from the given column.
func (d *Dataframe) GetColumnUnique(col interface{}) ([]string, error) {
	ret := simpleset.NewStringSet()
	c, err := d.getIndex(d.Header, d.ncol, "Header", col)
	if err == nil {
		for _, i := range d.Rows {
			ret.Add(i[c])
		}
	}
	return ret.ToStringSlice(), err
}
