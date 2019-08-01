[![Build Status](https://travis-ci.com/icwells/go-tools.svg?branch=master)](https://travis-ci.com/icwells/go-tools)

# go-tools  

## Commonly used GO functions (intended for personal use, but feel free to use)  

### To download and install to gopath:  
	go get github.com/icwells/go-tools/iotools  
	go get github.com/icwells/go-tools/strarray  

## strarray contains functions for working with slices and maps of strings  

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

## iotools wraps common file/path functions with error handling  

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

#### WriteToCSV
	iotools.WriteToCSV(outfile, header string, results [][]string)  

Writes header and slice of string slices to comma seperated file.  
Prints error if line cannot be written but does not exit.  
