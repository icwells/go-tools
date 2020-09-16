// Contains helper functions for i/o operations

package iotools

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

var DELIM = []string{"\t", ",", "|", ";", ":", " "}

// CheckError returns true if err is nil; otherwise, if code is 0 it prints a warning and returns false.
// If code is not 0, it prints an error formatted with msg and exits with code.
func CheckError(msg string, err error, code int) bool {
	ret := false
	if err == nil {
		ret = true
	} else if code == 0 {
		fmt.Fprintf(os.Stderr, "\t[Warning] %s: %v\n", msg, err)
	} else {
		fmt.Fprintf(os.Stderr, "\n\t[ERROR] %s: %v\n\n", msg, err)
		os.Exit(code)
	}
	return ret
}

// Exists seturns true if the given file or directory exists.
func Exists(path string) bool {
	ret := true
	_, err := os.Stat(path)
	if err != nil {
		ret = false
	}
	return ret
}

// GetGOPATH returns gopath environent environment variable.
func GetGOPATH() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	if Exists(gopath) == false {
		fmt.Print("\n\t[Error] Cannot determine GOPATH. Exiting.\n\n")
		os.Exit(10)
	}
	return gopath
}

// FormatPath returns path name with trailing slash (os.PathSeparator) and makes the directory if makenew is true.
func FormatPath(path string, makenew bool) (string, bool) {
	if path[len(path)-1] != os.PathSeparator {
		path = path + string(os.PathSeparator)
	}
	ex := Exists(path)
	if makenew == true {
		if ex == false {
			err := os.MkdirAll(path, os.ModePerm)
			if err == nil {
				// Change value of ex if mkdir was successful
				ex = true
			}
		}
	}
	return path, ex
}

// GetExt returns the file extension (everthing after the last period) from file.
func GetExt(file string) string {
	return strings.Replace(filepath.Ext(file), ".", "", 1)
}

// GetFileName returns the base name (everything between the final slash and first period) from file.
func GetFileName(file string) string {
	ret := ""
	sep := string(os.PathSeparator)
	if strings.Contains(file, ".") == true && strings.Contains(file, sep) == true {
		// Index slash first in case there is a period in path
		tmp := file[strings.LastIndex(file, sep)+1:]
		idx := strings.Index(tmp, ".")
		if idx >= 0 {
			ret = tmp[:idx]
		}
	}
	return ret
}

// GetParent returns name of parent directory from path.
func GetParent(path string) string {
	if strings.Contains(path, ".") {
		// Drop file name
		path = filepath.Dir(path)
	}
	if path[len(path)-1] == os.PathSeparator {
		// Drop trailing slash
		path = path[:len(path)-1]
	}
	idx := strings.LastIndex(path, string(os.PathSeparator)) + 1
	return path[idx:]
}

// GetHeader returns a map of header names to indeces.
func GetHeader(row []string) map[string]int {
	ret := make(map[string]int)
	for idx, i := range row {
		ret[i] = idx
	}
	return ret
}

// FindDelim returns delimiter (tab, comma, pipe, semicolon, colon, or space) from a text file. Returns an error if delimiter cannot be found.
func FindDelim(infile string) (string, error) {
	var count int
	var lines []string
	// Get first five lines
	f := OpenFile(infile)
	defer f.Close()
	input := GetScanner(f)
	for input.Scan() {
		count++
		lines = append(lines, strings.TrimSpace(string(input.Text())))
		if count == 5 {
			break
		}
	}
	for _, d := range DELIM {
		match := true
		count = strings.Count(lines[0], d)
		if count > 0 {
			for _, i := range lines[1:] {
				if strings.Count(i, d) != count {
					match = false
					break
				}
			}
			if match {
				return d, nil
			}
		}
	}
	return "", fmt.Errorf("[Warning] Cannot determine delimeter.")
}

// GetDelim returns delimiter (tab, comma, pipe, semicolon, colon, or space) from a line of a text file. Returns an error if delimiter cannot be found.
func GetDelim(header string) (string, error) {
	var d string
	var err error
	var max int
	for _, i := range DELIM {
		count := strings.Count(header, i)
		if count > max {
			d = i
			max = count
		}
	}
	if d == "" {
		err = fmt.Errorf("[Warning] Cannot determine delimeter.")
	}
	return d, err
}
