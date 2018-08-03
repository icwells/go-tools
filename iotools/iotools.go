// Handles basic file operations and handles errors.

package iotools

import (
	"fmt"
	"os"
	"strings"
)

func CheckError(msg string, err error, code int) bool {
	// If err is nil, returns true; otherwise, if code is 0 print warning and return false.
	// if code is not 0, print error and exit with code
	ret = false
	if err == nil {
		ret = true
	} else if code == 0 {
		fmt.Fprintf(os.Stderr, "\t[Warning] %s: %v\n", outfile, err)
	} else {
		fmt.Fprintf(os.Stderr, "\n\t[ERROR] %s: %v\n\n", outfile, err)
		os.Exit(code)
	}
	return ret
}

func OpenFile(file string) *os.File {
	// Returns file stream, exits if it encounters an error
	f, err := os.Open(file)
	_ := CheckError(fmt.Sprintf("Reading %s", file), err, 1)
	return f
}

func CreateFile(file string) *os.File {
	// Creates file and returns file stream, exits if it encounters an error
	f, err := os.Create(file)
	_ := CheckError(fmt.Sprintf("Creating %s", file), err, 2)
	return f
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
	idx := strings.LastIndex(file, ".") + 1
	return file[idx:]
}

func GetFileName(file string) string {
	// Returns base name from filename
	idx := strings.LastIndex(file, "/") + 1
	ind := strings.Index(file, ".")
	return file[idx:ind]
}

func GetParent(path string) string {
	// Returns name of parent directory from filename/directory
	if strings.Contains(path, ".") == true && path[len(path)-1] == '/' {
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

func WriteToCSV(outfile, header string, results [][]string) {
	// Writes slice of slices to file
	out := CreateFile(outfile)
	defer out.Close()
	_, err := out.WriteString(header + "\n")
	_ := CheckError(fmt.Sprintf("Writing header to %s", outfile), err, 0)
	for _, i := range results {
		// Write comma seperated items to file
		_, err = out.WriteString(strings.Join(i, ",") + "\n")
		_ := CheckError(fmt.Sprintf("Writing header to %s", outfile), err, 0)
	}
}
