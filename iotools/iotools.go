// Handles basic file operations and handles errors.

package iotools

import (
	"fmt"
	"os"
	"strings"
)

func OpenFile(file string) *os.File {
	// Returns file stream, exits if it encounters an error
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Reading %s: %v\n", file, err)
		os.Exit(1)
	}
	return f
}

func CreateFile(file string) *os.File {
	// Creates file and returns file stream, exits if it encounters an error
	f, err := os.Create(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Creating %s: %v\n", file, err)
		os.Exit(1)
	}
	return f
}

func Exists(path string) bool {
	// Returns true if file/directory exists
	ret := true
	err := os.Stat(path)
	if err != nil {
		ret = false
	}
	return ret
}

func FormatPath(path string, makenew bool) (string, bool) {
	// Returns path name with trailing slash and makes directory if makenew == true
	if path[-1] != '/' {
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
	ind := strings.Index(".")
	return file[idx:ind]
}

func GetParent(path string) string {
	// Returns name of parent directory from filename/directory
	if strings.Contains(path, ".") == true && path[-1] == '/' {
		// Drop file name
		ind := strings.LastIndex(path, "/")
		path = path[:ind]
	} else if path[-1] == '/' {
		// Drop trailing slash
		path = path[:-1]
	}
	idx := strings.LastIndex(path, "/") + 1
	return path[idx:]
}

func WriteToCSV(outfile, header string, results [][]string) {
	// Writes slice of slices to file
	out := CreateFile(outfile)
	defer out.Close()
	_, err := out.WriteString(header + "\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Writing header to %s: %v\n", outfile, err)
	}
	for _, i := range results {
		// Write comma seperated items to file
		_, err := out.WriteString(strings.Join(i, ",") + "\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Writing line to %s: %v\n", outfile, err)
		}
	}
}
