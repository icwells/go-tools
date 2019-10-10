[![Build Status](https://travis-ci.com/icwells/go-tools.svg?branch=master)](https://travis-ci.com/icwells/go-tools)

# go-tools  

## Commonly used GO functions (intended for personal use, but feel free to use)  

Copyright 2019 by Shawn Rupp

1. [Description](#Description)
2. [Installation](#Installation)  
3. [iotools](#iotools)  
4. [strarray](#strarray)  
5. [Set](#Set)  
6. [dataframe](#dataframe)

## Description  

## Installation  
	go get github.com/icwells/go-tools/dataframe  
	go get github.com/icwells/go-tools/iotools  
	go get github.com/icwells/go-tools/strarray  

## iotools  
Wraps common file/path functions with error handling and provides basic input/output functions.  

#### CheckError  
	iotools.CheckError(msg string, err error, code int) bool  

Returns true if err is nil. If err is not nil and code is 0, it prints a warning formatted with 
msg and returns false. Otherwise, it prints and error formated with message and exits with code.  

#### OpenFile
	iotools.OpenFile(file string) *os.File  

Returns pointer to File type. Prints an error and exits if file cannot be opened.  

#### CreateFile
	iotools.CreateFile(file string) *os.File   

Creates new file and returns pointer. Prints an error and exits if file cannot be created.

#### AppendFile
	iotools.AppendFile(file string) *os.File   

Returns pointer to file to append to. Creates file if it does not exist. 
Prints an error and exits if file cannot be opened or created.

#### GetScanner  
	iotools.GetScanner(f *os.File) *bufio.Scanner  

Returns scanner for either gzipped or uncompressed file.  

#### Exists  
	iotools.Exists(path string) bool

Returns true if file or directory exists. Otherwise, returns false.  

#### GetGOPATH  
	iotools.getGOPATH() string

Returns GOPATH from environment. Prints error and exits if it cannot detemermine GOPATH.  

#### FormatPath  
	iotools.FormatPath(path string, makenew bool) (string, bool)  

Returns path name with trailing slash and result of Exists(path). Makes directory if makenew == true.  

#### GetExt  
	iotools.GetExt(file string) string  

Returns extension from file name.  

#### GetFileName  
	iotools.GetFileName(file string) string  

Returns base name from file name.  

#### GetParent  
	iotools.GetParent(file string) string  

Returns name of parent directory from file or directory.  

#### GetHeader
	iotools.GetHeader(header []string) map[string]int  

Returns map of column names matched to index numbers for simple header parsing.  

#### GetDelim
	iotools.GetDelim(header string) string  

Returns delimiter from header of a text file.  

#### ReadFile  
	iotools.ReadFile(infile string, header bool) ([][]string, map[string]int)  

Reads tab/comma/space delimiter text files, trims newlines, and returns header as map of 
column name: index and data as slice of rows split by delimiter.  

#### WriteToCSV
	iotools.WriteToCSV(outfile, header string, results [][]string)  

Writes header and slice of string slices to comma seperated file.  
Prints error if line cannot be written but does not exit.  

## strarray  
Contains functions for working with slices and maps of strings, as well as a Python-style set.   

#### TitleCase(t string) string  
Manually converts term to title case (strings.Title is buggy).  

#### InSliceStr
	strarray.InSliceStr(l []string, s string) bool  

Returns true if s is in l.  

#### InSliceSli
	strarray.InSliceSli(l [][]string, s string, c int) bool  

Returns true if s is in column c in l  

#### SliceIndex  
	strarray.SliceIndex(l []string, v string) int  

Returns first index v value in slice. Returns -1 if it is not found.  

#### SliceCount  
	strarray.SliceCount(s []string, v string) int  

Returns number of occurances of v in s.  

#### DeleteSliceIndex  
	strarray.DeleteSliceIndex(s []string, idx int) []string  

Deletes item at idx while preventing index errors.  

##### DeleteSliceValue  
	strarray.DeleteSliceValue(s []string, v string) []string  

Deletes all occurances of v from s.  

### Set  
The set struct is a simple python-style set for strings.  

#### NewSet  
	strarray.NewSet() Set

Initializes new set.  

#### ToSet
	strarray.ToSet(s []string) Set  

Converts slice of strings to set.  

#### Length  
	set.Length()  

Returns length of set.  

#### Add  
	set.Add(value string)  

Adds string value to set.  

#### Extend  
	set.Extend(v []string)  

Adds all elements of slice to set.  

#### Pop  
	set.Pop(v string)  

Removes v from set.  

#### InSet  
	set.InSet(value string)  

Reurns true if value is in the set. Returns false if it is not.  

#### ToSlice  
	set.ToSlice() []string  

Returns set as a sorted string slice.

## dataframe  
Provides a variable length, two-dimensional array of strings which can be indexed by row/column names 
or numbers. It is meant to quickly and cleanly parse input data, particuly when the data of interest contains text. 
This is still in developement, and there are more features to come.  

### The Dataframe Struct  
The Dataframe struct stores tabular data in a two-dimensional slice of strings. It stores a header as a map with string keys 
and column indeces as values. It will optionally also store an index containing string identifiers with row indeces as values.  

	Rows   [][]string
	Header map[string]int
	Index  map[string]int

There are two ways to make a new dataframe. One is to initialize an empty struct, while the other is to read input data 
directly from a file. The columnn value indicates which column should be used for the row index. A negative value 
will omit the index (note that sting indeces cannot be used if there is not index).  

#### dataframe.NewDataFrame(column int) *Dataframe  
Initializes an empty dataframe. The given column number of any input data will be used as the index column (a negative value will omit the index).  

#### DataFrameFromFile(infile string, column int) *Dataframe  
Creates a dataframe the same as above, but loads in data from the given input file. the first row is assumed to be the header.  

### Setter Functions  

#### Dataframe.AddRow(row []string) error  
Adds row to dataframe. If using an index, the index column will be subset from the slice and added to the index. Returns an error if the index value is already present.  

#### Dataframe.SetHeader(row []string)  
Converts given row to header. Subsets index column if using an index.  

### Getter Functions  

### Other 

#### Dataframe) ToCSV(outfile string)  
Writes rows to csv with index and header. Includes header value for index column.  

#### Dataframe.DeleteRow(idx interface{}) error  
Deletes given row from dataframe and adjusts index (if using). Returns error if given index is not found.  

#### Dataframe.DeleteColumn(col interface{}) error  
Deletes given column from dataframe and adjusts index (if using). Returns error if given index is not found.  
