// Handles basic file operations and handles errors.

package iotools

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

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

// chekcFile exists if file does not exist.
func checkFile(infile string) {
	if !Exists(infile) {
		fmt.Printf("\n\t[Error] Input file %s not found. Exiting.\n\n", infile)
		os.Exit(1)
	}
}

// setHeader returns the header and delimiter from the input file.
func setHeader(infile string, header bool) (map[string]int, string) {
	var h map[string]int
	d, err := FindDelim(infile)
	if err != nil {
		fmt.Println(err)
	}
	if header {
		f := OpenFile(infile)
		defer f.Close()
		input := GetScanner(f)
		for input.Scan() {
			// Get header and delimiter from first line
			line := strings.TrimSpace(string(input.Text()))
			h = GetHeader(strings.Split(line, d))
			break
		}
	}
	return h, d
}

// YieldFile reads compressed and uncompressed text files and returns the header as a map of indeces and the lines as an iterable channel of string slices.
//		reader, header := iotools.YieldFile(infile, true)
//      for i := range reader { ...
func YieldFile(infile string, header bool) (<-chan []string, map[string]int) {
	var h map[string]int
	ch := make(chan []string)
	checkFile(infile)
	h, d := setHeader(infile, header)
	go func() {
		f := OpenFile(infile)
		defer f.Close()
		input := GetScanner(f)
		for input.Scan() {
			if !header {
				var s []string
				line := strings.TrimSpace(string(input.Text()))
				if d == "" {
					s = append(s, line)
				} else {
					s = strings.Split(line, d)
				}
				for idx, i := range s {
					s[idx] = strings.TrimSpace(i)
				}
				ch <- s
			} else {
				// Skip first line
				header = false
			}
		}
		close(ch)
	}()
	return ch, h
}

// ReadFile reads in compressed and uncompressed text files as a two dimensional slice of strings and the header as a map of indeces.
func ReadFile(infile string, header bool) ([][]string, map[string]int) {
	var ret [][]string
	reader, h := YieldFile(infile, header)
	for i := range reader {
		ret = append(ret, i)
	}
	return ret, h
}

// Gzip compresses infile to same directory. Adds '.gz' extension.
func Gzip(outfile, header string, results [][]string) {
	f := CreateFile(outfile)
	defer f.Close()
	w := gzip.NewWriter(f)
	defer w.Close()
	// Write bytes in compressed form to the file.
	w.Write([]byte(header))
	for _, i := range results {
		w.Write([]byte(strings.Join(i, ",") + "\n"))
	}
}

// WriteToCSV writes a string header and a two dimensional slice of strings to csv. Gzips file if last three characters of file name are '.gz'.
func WriteToCSV(outfile, header string, results [][]string) {
	if header[len(header)-1] != '\n' {
		header += "\n"
	}
	if strings.Contains(outfile, ".gz") {
		Gzip(outfile, header, results)
	} else {
		out := CreateFile(outfile)
		defer out.Close()
		_, err := out.WriteString(header)
		_ = CheckError(fmt.Sprintf("Writing header to %s", outfile), err, 0)
		for _, i := range results {
			// Write comma seperated items to file
			_, err = out.WriteString(strings.Join(i, ",") + "\n")
			_ = CheckError(fmt.Sprintf("Writing results to %s", outfile), err, 0)
		}
	}
}
