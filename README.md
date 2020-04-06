[![Build Status](https://travis-ci.com/icwells/go-tools.svg?branch=master)](https://travis-ci.com/icwells/go-tools)

# go-tools  

## Commonly used GO functions (intended for personal use, but feel free to use)  

Copyright 2019 by Shawn Rupp

1. [dataframe](#dataframe)    
2. [fraction](#fraction)  
3. [iotools](#iotools)  
4. [strarray](#strarray)  

## Note: strarray.Set is deprecated.  
It can be replaced by [simpleset](https://github.com/icwells/simpleset). Simply install simpleset:  

	go get github.com/icwells/simpleset   

Then change the import path to "github.com/icwells/simpleset" and the call the constructor function to simpleset.NewStringSet() 
(Note that this will return a pointer).  

## dataframe  
Provides a variable length, two-dimensional array of strings which can be indexed by row/column names 
or numbers. It is meant to quickly and cleanly parse input data, particuly when the data of interest contains text. 

	go get github.com/icwells/go-tools/dataframe  

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

#### dataframe.FromFile(infile string, column interface{}) (*Dataframe, error)  
Creates a dataframe the same as above, but loads in data from the given input file. The first row is assumed to be the header.  

[![GoDoc](https://godoc.org/github.com/icwells/go-tools/dataframe?status.svg)](https://godoc.org/github.com/icwells/go-tools/dataframe)

## fraction  
Provides a struct to store fractions and provides mathmatical and conversion methods.  

	go get github.com/icwells/go-tools/fraction  

## iotools  
Wraps common file/path functions with error handling and provides basic input/output functions.  

	go get github.com/icwells/go-tools/iotools  

[![GoDoc](https://godoc.org/github.com/icwells/go-tools/iotools?status.svg)](https://godoc.org/github.com/icwells/go-tools/iotools)

## strarray  
Contains functions for working with slices and maps of strings, as well as a Python-style set.   

	go get github.com/icwells/go-tools/strarray  

[![GoDoc](https://godoc.org/github.com/icwells/go-tools/strarray?status.svg)](https://godoc.org/github.com/icwells/go-tools/strarray)
