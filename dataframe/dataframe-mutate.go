// Getter methods for dataframe

package dataframe

import (
	"errors"
	"fmt"
)

// RenameColumn changes the name of column o to n. Returns an error if o is not found.
func (d *Dataframe) RenameColumn(o, n string) error {
	var err error
	if len(n) == 0 {
		err = errors.New("New column name is empty.")
	} else if _, ex := d.Header[n]; ex {
		err = fmt.Errorf("Column %s is already present in header.", n)
	} else if v, ex := d.Header[o]; ex {
		delete(d.Header, o)
		d.Header[n] = v
	} else {
		err = fmt.Errorf("Column %s was not found in header.", o)
	}
	return err
}

// Compare returns an error if target dataframe is not equal to d (for testing).
func (d *Dataframe) Compare(n *Dataframe) error {
	// Compares output of equivalent tables
	var err error
	nc, nr := n.Dimensions()
	dc, dr := d.Dimensions()
	if nc != dc && nr != dr {
		err = fmt.Errorf("Target dataframe dimensions [%d, %d] do not equal original: [%d, %d]", nc, nr, dc, dr)
	} else {
		for key := range n.Index {
			for k := range n.Header {
				var a, e string
				if a, err = n.GetCell(key, k); err == nil {
					e, err = d.GetCell(key, k)
					if err == nil && a != e {
						// Make sure error is not due to floating point precision
						af, aerr := n.GetCellFloat(key, k)
						ef, eerr := n.GetCellFloat(key, k)
						if eerr != nil || aerr != nil || af != ef {
							err = fmt.Errorf("%s-%s: Actual value %s does not equal expected: %s", key, k, a, e)
						}
					}
				}
			}
		}
	}
	return err
}

// Extend appends rows from n to the current dataframe. Rows with redundant index values will be skipped.
func (d *Dataframe) Extend(n *Dataframe) error {
	if d.ncol != n.ncol {
		return fmt.Errorf("New dataframe width %d does not equal width of parent database %d.", n.ncol, d.ncol)
	}
	if n.col >= 0 {
		for k := range n.Index {
			row, _ := n.GetRow(k)
			d.appendRow(k, row)
		}
	} else {
		for _, i := range n.Rows {
			d.appendRow("", i)
		}
	}
	return nil
}

// Deletes given entry from index/header and decrements higher values
func (d *Dataframe) decrememntMap(m map[string]int, n int) map[string]int {
	for k, v := range m {
		if v > n {
			m[k]--
		} else if v == n {
			delete(m, k)
		}
	}
	return m
}

// DeleteRow removes the given row from dataframe.
func (d *Dataframe) DeleteRow(idx interface{}) error {
	r, err := d.getIndex(d.Index, d.nrow, "Index", idx)
	if err == nil {
		// Remove row and decrement counter
		if r == 0 {
			d.Rows = d.Rows[1:]
		} else if r == d.nrow-1 {
			d.Rows = d.Rows[:r]
		} else {
			d.Rows = append(d.Rows[:r], d.Rows[r+1:]...)
		}
		d.nrow--
		if d.col >= 0 {
			d.Index = d.decrememntMap(d.Index, r)
		}
	}
	return err
}

// DeleteColumn removes the given column from dataframe.
func (d *Dataframe) DeleteColumn(col interface{}) error {
	c, err := d.getIndex(d.Header, d.ncol, "Header", col)
	if err == nil {
		// Remove row and decrement counter
		for idx, i := range d.Rows {
			if c == 0 {
				d.Rows[idx] = i[1:]
			} else if c == d.nrow-1 {
				d.Rows[idx] = i[:c]
			} else {
				d.Rows[idx] = append(i[:c], i[c+1:]...)
			}
		}
		d.ncol--
		d.Header = d.decrememntMap(d.Header, c)
	}
	return err
}
