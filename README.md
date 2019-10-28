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

#### iotools.CheckError(msg string, err error, code int) bool  
Returns true if err is nil. If err is not nil and code is 0, it prints a warning formatted with 
msg and returns false. Otherwise, it prints and error formated with message and exits with code.  

#### iotools.OpenFile(file string) *os.File  
Returns pointer to File type. Prints an error and exits if file cannot be opened.  

#### iotools.CreateFile(file string) *os.File   
Creates new file and returns pointer. Prints an error and exits if file cannot be created.

#### iotools.AppendFile(file string) *os.File   
Returns pointer to file to append to. Creates file if it does not exist. 
Prints an error and exits if file cannot be opened or created.

#### iotools.GetScanner(f *os.File) *bufio.Scanner  
Returns scanner for either gzipped or uncompressed file.  

#### iotools.Exists(path string) bool
Returns true if file or directory exists. Otherwise, returns false.  

#### iotools.getGOPATH() string
Returns GOPATH from environment. Prints error and exits if it cannot detemermine GOPATH.  

#### iotools.FormatPath(path string, makenew bool) (string, bool)  
Returns path name with trailing slash and result of Exists(path). Makes directory if makenew == true.  

#### iotools.GetExt(file string) string  
Returns extension from file name.  

#### iotools.GetFileName(file string) string  
Returns base name from file name.  

#### iotools.GetParent(file string) string  
Returns name of parent directory from file or directory.  

#### iotools.GetHeader(header []string) map[string]int  
Returns map of column names matched to index numbers for simple header parsing.  

#### iotools.GetDelim(header string) string  
Returns delimiter from header of a text file.  

#### iotools.ReadFile(infile string, header bool) ([][]string, map[string]int)  
Reads tab/comma/space delimiter text files, trims newlines, and returns header as map of 
column name: index and data as slice of rows split by delimiter.  

#### iotools.WriteToCSV(outfile, header string, results [][]string)  
Writes header and slice of string slices to comma seperated file.  
Prints error if line cannot be written but does not exit.  

## strarray  
Contains functions for working with slices and maps of strings, as well as a Python-style set.   

#### TitleCase(t string) string  
Manually converts term to title case (strings.Title is buggy).  

#### strarray.InSliceStr(l []string, s string) bool  
Returns true if s is in l.  

#### strarray.InSliceSli(l [][]string, s string, c int) bool  
Returns true if s is in column c in l  

#### strarray.SliceIndex(l []string, v string) int  
Returns first index v value in slice. Returns -1 if it is not found.  

#### strarray.SliceCount(s []string, v string) int  
Returns number of occurances of v in s.  

#### strarray.DeleteSliceIndex(s []string, idx int) []string  
Deletes item at idx while preventing index errors.  

##### strarray.DeleteSliceValue(s []string, v string) []string  
Deletes all occurances of v from s.  

### Set  
The set struct is a simple python-style set for strings.  

#### strarray.NewSet() Set
Initializes new set.  

#### strarray.ToSet(s []string) Set  
Converts slice of strings to set.  

#### set.Length()  
Returns length of set.  

#### set.Add(value string)  
Adds string value to set.  

#### set.Extend(v []string)  
Adds all elements of slice to set.  

#### set.Pop(v string)  
Removes v from set.  

#### set.InSet(value string)  
Reurns true if value is in the set. Returns false if it is not.  

#### set.ToSlice() []string  
Returns set as a sorted string slice.

## dataframe  
Provides a variable length, two-dimensional array of strings which can be indexed by row/column names 
or numbers. It is meant to quickly and cleanly parse input data, particuly when the data of interest contains text. 
This is still in developement, and there are more features to come.  

### The Dataframe Struct  
The Dataframe struct stores tabular data in a two-dimensional slice of strings. It stores a header as a map with string keys 
and column indeces as values. It will optionally also store an index containing string identifiers with row indeces as values. 
The data (in the Rows slice), header, and index are all exported and can be directly modified for greater flexibility 
(although this can lead to errors).  

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

Dataframe.Extend(df *Dataframe)  
Adds rows from new dataframe to existing Rows slice.  

### Getter Functions  

#### Dataframe.Length() int
Returns number of rows.  

#### Dataframe.Dimensions() (int, int)
Returns number of columns and number of rows respectively.  

#### Dataframe.GetHeader() []string  
Returns header as string slice.  

#### Dataframe.GetIndex() []string  
Returns index as string slice.  

#### Dataframe.UpdateCell(idx interface{}, col interface{}, v string) error  
Replaces given cell with value. Returns and error if idx an col are not a string or int.  

#### Dataframe.GetCell(idx interface{}, col interface{}) (string, error)  
Returns given cell from dataframe as a string. Returns an error if idx and col are not a string or int.  

#### Dataframe.GetCellInt(idx interface{}, col interface{}) (int, error)  
Returns given cell as an integer. Returns an error if idx and col are not a string or int.  

#### Dataframe.GetCellFloat(idx interface{}, col interface{}) (float64, error)  
Returns given cell as float64. Returns an error if idx and col are not a string or int.  

#### Dataframe.GetRow(idx interface{}) ([]string, error)  
Returns given row from dataframe. Returns an error if idx is not a string or int.  

#### Dataframe.GetColumn(col interface{}) ([]string, error)  
Returns given column from dataframe. Returns an error if col is not a string or int.  

#### Dataframe.GetColumnUnique(col interface{}) ([]string, error)  
Returns unique values from given column. Returns an error if col is not a string or int.  

### Other 

#### Dataframe) ToCSV(outfile string)  
Writes rows to csv with index and header. Includes header value for index column.  

#### Dataframe.DeleteRow(idx interface{}) error  
Deletes given row from dataframe and adjusts index (if using). Returns error if given index is not found.  

#### Dataframe.DeleteColumn(col interface{}) error  
Deletes given column from dataframe and adjusts index (if using). Returns error if given index is not found.  
