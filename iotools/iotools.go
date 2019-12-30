// Handles basic file operations and handles errors.

package iotools

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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

// OpenFile opens the file at the given file path and returns a file stream.
// It will exit if it encounters an error.
func OpenFile(file string) *os.File {
	f, err := os.Open(file)
	_ = CheckError(fmt.Sprintf("Reading %s", file), err, 1)
	return f
}

// CreateFile creates a file at the given file path and returns a file stream.
// It will exit if it encounters an error.
func CreateFile(file string) *os.File {
	// Creates file and returns file stream, exits if it encounters an error
	f, err := os.Create(file)
	_ = CheckError(fmt.Sprintf("Creating %s", file), err, 2)
	return f
}

// AppendFile opens the file at the given file path and returns file stream in append mode.
// It will exit if it encounters an error.
func AppendFile(file string) *os.File {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_ = CheckError(fmt.Sprintf("Append to %s", file), err, 3)
	return f
}

// GetScanner determines if a file is gzipped or not returns the appropriate scanner.
// It will exit if it encounters an error.
func GetScanner(f *os.File) *bufio.Scanner {
	var scanner *bufio.Scanner
	reader := bufio.NewReader(io.Reader(f))
	// Check if first two bytes == 0x1f8b (i.e. 31 & 139)
	test, err := reader.Peek(2)
	_ = CheckError("Cannot peek into file", err, 3)
	if test[0] == 31 && test[1] == 139 {
		// Make scanner from gzip reader
		greader, _ := gzip.NewReader(reader)
		_ = CheckError("Cannot read gzipped file", err, 4)
		scanner = bufio.NewScanner(greader)
	} else {
		// Make scanner from bufio reader
		scanner = bufio.NewScanner(reader)
	}
	return scanner
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

// GetDelim returns delimiter (tab, comma, or space) from a text file.
func GetDelim(header string) string {
	var d string
	found := false
	for _, i := range []string{"\t", ",", " "} {
		if strings.Contains(header, i) == true {
			d = i
			found = true
			break
		}
	}
	if found == false {
		fmt.Print("\n\t[Error] Cannot determine delimeter. Exiting.\n\n")
	}
	return d
}

// ReadFile reads in compressed and uncompressed text files as a two dimensional slice of strings and the header as a map of indeces.
func ReadFile(infile string, header bool) ([][]string, map[string]int) {
	var d string
	var h map[string]int
	var ret [][]string
	if !Exists(infile) {
		fmt.Printf("\n\t[Error] Input file %s not found. Exiting.\n\n", infile)
		os.Exit(1)
	}
	f := OpenFile(infile)
	defer f.Close()
	input := GetScanner(f)
	for input.Scan() {
		line := strings.TrimSpace(string(input.Text()))
		if d == "" {
			d = GetDelim(line)
		}
		s := strings.Split(line, d)
		if !header {
			for idx, i := range s {
				s[idx] = strings.TrimSpace(i)
			}
			ret = append(ret, s)
		} else {
			h = GetHeader(s)
			header = false
		}
	}
	return ret, h
}

// WriteToCSV writes a string header and a two dimensional slice of strings to csv.
func WriteToCSV(outfile, header string, results [][]string) {
	out := CreateFile(outfile)
	defer out.Close()
	_, err := out.WriteString(header + "\n")
	_ = CheckError(fmt.Sprintf("Writing header to %s", outfile), err, 0)
	for _, i := range results {
		// Write comma seperated items to file
		_, err = out.WriteString(strings.Join(i, ",") + "\n")
		_ = CheckError(fmt.Sprintf("Writing header to %s", outfile), err, 0)
	}
}
