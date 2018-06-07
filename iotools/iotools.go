// Handles basic file operations file open/close and handles errors.

package iotools

import (
	"fmt"
	"os"
	"strings"
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
		fmt.Fprintf(os.Stderr, "[ERROR] Creating %s: %v\n", file, err)
		os.Exit(1)
	}
	return f
}

func writeToCSV(outfile, header string, results []string) {
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
	}
}
