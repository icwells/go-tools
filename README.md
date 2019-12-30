[![Build Status](https://travis-ci.com/icwells/go-tools.svg?branch=master)](https://travis-ci.com/icwells/go-tools)

# go-tools  

## Commonly used GO functions (intended for personal use, but feel free to use)  

Copyright 2019 by Shawn Rupp

1. [Installation](#Installation)  
2. [iotools](#iotools)  
3. [strarray](#strarray)  
4. [dataframe](#dataframe)

## Installation  
	go get github.com/icwells/go-tools/dataframe  
	go get github.com/icwells/go-tools/iotools  
	go get github.com/icwells/go-tools/strarray  

## Note: strarray.Set is deprecated.  
It can be replaced by [simpleset](https://github.com/icwells/simpleset). Simply install simpleset:  

	go get github.com/icwells/simpleset   

Then change the import path to "github.com/icwells/simpleset" and the call the constructor function to simpleset.NewStringSet() 
(Note that this will return a pointer).  

## iotools  
Wraps common file/path functions with error handling and provides basic input/output functions.  

See GoDocs link for usage.

## strarray  
Contains functions for working with slices and maps of strings, as well as a Python-style set.   

See GoDocs link for usage.

## dataframe  
Provides a variable length, two-dimensional array of strings which can be indexed by row/column names 
or numbers. It is meant to quickly and cleanly parse input data, particuly when the data of interest contains text. 

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

#### dataframe.NewDataFrame(column interface{}) (*Dataframe, error)  
Initializes an empty dataframe. The given column name/number (must be string or int) of any input data will be used as the index column (a negative value will omit the index).  

#### dataframe.DataFrameFromFile(infile string, column interface{}) (*Dataframe, error)  
Creates a dataframe the same as above, but loads in data from the given input file. The first row is assumed to be the header.  

### Setter Functions  

#### Dataframe.AddRow(row []string) error  
Adds row to dataframe. If using an index, the index column will be subset from the slice and added to the index. Returns an error if the index value is already present.  

#### Dataframe.SetHeader(row []string) error  
Converts given row to header. Subsets index column if using an index.  

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

#### Dataframe.SliceRow(idx interface{}, start interface{}, end interface{}) ([]string, error)  
Returns subset of row (idx) between start and end (row[start:end]). A negative value for end will return all values after start (row[start:]).  

#### Dataframe.GetColumn(col interface{}) ([]string, error)  
Returns given column from dataframe. Returns an error if col is not a string or int.  

#### Dataframe.GetColumnUnique(col interface{}) ([]string, error)  
Returns unique values from given column. Returns an error if col is not a string or int.  

### Other 

#### Dataframe.ToSlice() [][]string  
Returns dataframe as slice of string slices (inserts index values if needed).  

#### Dataframe.ToCSV(outfile string)  
Writes rows to csv with index and header. Includes header value for index column.  

#### Dataframe.DeleteRow(idx interface{}) error  
Deletes given row from dataframe and adjusts index (if using). Returns error if given index is not found.  

#### Dataframe.DeleteColumn(col interface{}) error  
Deletes given column from dataframe and adjusts index (if using). Returns error if given index is not found.  
