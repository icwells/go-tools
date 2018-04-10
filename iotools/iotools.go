// Handles basic file operations file open/close and handles errors.

package iotools

import (
	"fmt"
	"os"
)

func OpenFile(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Reading %s: %v\n", file, err)
		os.Exit(1)
	}
	return f
}

func CreateFile(file string) *os.File {
	f, err := os.Create(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Reading %s: %v\n", file, err)
		os.Exit(1)
	}
	return f
}
