// Commonly used functions for handling string-based arrays

package strarray

func InSliceStr(l []string, s string) bool {
	// Returns true if s is in l
	in := false
	for _, i := range l {
		if s == i {
			in = true
			break
		}
	}
	return in
}

func InSliceSli(l [][]string, s string, c int) bool {
	// Returns true if s is in column c in l
	in := false
	for _, i := range l {
		if s == i[c] {
			in = true
			break
		}
	}
	return in
}

func InMapStr(m map[string]string, s string) bool {
	// Returns true if s is a key in m
	in := false
	for k, _ := range m {
		if s == k {
			in = true
			break
		}
	}
	return in
}

func InMapSli(m map[string][]string, s string) bool {
	// Returns true if s is a key in m
	in := false
	for k, _ := range m {
		if s == k {
			in = true
			break
		}
	}
	return in
}
