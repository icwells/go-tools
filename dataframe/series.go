// Defines single row series struct and methods

package dataframe

import (
	"fmt"
	"strconv"
	"strings"
)

type Series struct {
	// Encapsulates error so it can be returned through channel
	Error error
	// Same as dataframe header
	Header map[string]int
	// Interger index of row in Rows slice
	Index int
	// Name of row if index map is set
	Name string
	// Copy of row from dataframe at index
	Row []string
}

// GetCell returns cell from the given column.
func (s *Series) GetCell(col string) (string, error) {
	var err error
	var ret string
	if v, ex := s.Header[col]; ex {
		ret = s.Row[v]
	} else {
		err = fmt.Errorf("Column name %s cannot be found in header", col)
	}
	return ret, err
}

// GetCellInt returns the value of the given cell as int. Returns an error if it cannot be converted.
func (s *Series) GetCellInt(col string) (int, error) {
	var ret int
	val, err := s.GetCell(col)
	if err == nil {
		// Remove any decimal values
		val = strings.Split(val, ".")[0]
		ret, err = strconv.Atoi(val)
	}
	return ret, err
}

// GetCellFloat returns the value of the given cell as float64. Returns an error if it cannot be converted.
func (s *Series) GetCellFloat(col string) (float64, error) {
	var ret float64
	val, err := s.GetCell(col)
	if err == nil {
		ret, err = strconv.ParseFloat(val, 64)
	}
	return ret, err
}

// ToSeries returns row at given index to a series
func (d *Dataframe) ToSeries(idx interface{}) *Series {
	var r int
	s := new(Series)
	r, s.Error = d.getIndex(d.Index, d.nrow, "Index", idx)
	if s.Error == nil {
		switch i := idx.(type) {
		case string:
			s.Name = string(i)
		}
		s.Header = d.Header
		s.Index = r
		s.Row = d.Rows[r]
	}
	return s
}

// Iterate returns each row with its index
func (d *Dataframe) Iterate() <-chan *Series {
	ch := make(chan *Series)
	go func() {
		if d.col >= 0 {
			for k := range d.Index {
				ch <- d.ToSeries(k)
			}
		} else {
			for idx := range d.Rows {
				ch <- d.ToSeries(idx)
			}
		}
		close(ch)
	}()
	return ch
}
