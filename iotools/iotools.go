// Handles basic file operations and handles errors.

package iotools

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"go/build"
	"io"
	"os"
	"strings"
)

func CheckError(msg string, err error, code int) bool {
	// If err is nil, returns true; otherwise, if code is 0 print warning and return false.
	// if code is not 0, print error and exit with code
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

func OpenFile(file string) *os.File {
	// Returns file stream, exits if it encounters an error
	f, err := os.Open(file)
	_ = CheckError(fmt.Sprintf("Reading %s", file), err, 1)
	return f
}

func CreateFile(file string) *os.File {
	// Creates file and returns file stream, exits if it encounters an error
	f, err := os.Create(file)
	_ = CheckError(fmt.Sprintf("Creating %s", file), err, 2)
	return f
}

func AppendFile(file string) *os.File {
	// Returns files stream to append to given file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_ = CheckError(fmt.Sprintf("Append to %s", file), err, 3)
	return f
}

func GetScanner(f *os.File) *bufio.Scanner {
	// Returns scanner for gzipped/uncompressed file
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

func Exists(path string) bool {
	// Returns true if file/directory exists
	ret := true
	_, err := os.Stat(path)
	if err != nil {
		ret = false
	}
	return ret
}

func GetGOPATH() string {
	// Returns gopath
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

func FormatPath(path string, makenew bool) (string, bool) {
	// Returns path name with trailing slash and makes directory if makenew == true
	if path[len(path)-1] != '/' {
		path = path + "/"
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

func GetExt(file string) string {
	// Returns extension from filename
	ret := ""
	if strings.Contains(file, ".") == true {
		idx := strings.LastIndex(file, ".") + 1
		ret = file[idx:]
	}
	return ret
}

func GetFileName(file string) string {
	// Returns base name from filename
	ret := ""
	if strings.Contains(file, ".") == true && strings.Contains(file, "/") == true {
		// Index slash first in case there is a period in path
		tmp := file[strings.LastIndex(file, "/")+1:]
		idx := strings.Index(tmp, ".")
		if idx >= 0 {
			ret = tmp[:idx]
		}
	}
	return ret
}

func GetParent(path string) string {
	// Returns name of parent directory from filename/directory
	if strings.Contains(path, ".") == true && path[len(path)-1] != '/' {
		// Drop file name
		ind := strings.LastIndex(path, "/")
		path = path[:ind]
	} else if path[len(path)-1] == '/' {
		// Drop trailing slash
		path = path[:len(path)-1]
	}
	idx := strings.LastIndex(path, "/") + 1
	return path[idx:]
}

func GetHeader(row []string) map[string]int {
	// Returns map of header indeces
	ret := make(map[string]int)
	for idx, i := range row {
		ret[i] = idx
	}
	return ret
}

func GetDelim(header string) string {
	// Returns delimiter
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

func ReadFile(infile string, header bool) ([][]string, map[string]int) {
	// Reads in text file as slice of string slices and header as map of indeces
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
		} else if header == false {
			ret = append(ret, strings.Split(line, d))
		} else {
			h = GetHeader(strings.Split(line, d))
			header = false
		}
	}
	return ret, h
}

func WriteToCSV(outfile, header string, results [][]string) {
	// Writes slice of slices to file
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
