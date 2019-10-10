// Getter methods for dataframe

package dataframe

func (d *Dataframe) decrememntMap(m map[string]int, n int) map[string]int {
	// Deletes given entry from index/header and decrements higher values
	for k, v := range m {
		if v > n {
			m[k]--
		} else if v == n {
			delete(m, k)
		}
	}
	return m
}

func (d *Dataframe) DeleteRow(idx interface{}) error {
	// Deletes given row
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

func (d *Dataframe) DeleteColumn(col interface{}) error {
	// Deletes given row
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
