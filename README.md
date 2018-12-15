# go-tools  

## Commonly used GO functions (intended for personal use, but feel free to use)  

### To download and install to gopath:  
	go get github.com/icwells/go-tools/iotools  
	go get github.com/icwells/go-tools/strarray  

## strarray contains functions for working with slices and maps of strings  

#### InSliceStr
	strarray.InSliceStr(l []string, s string) bool  

Returns true if s is in l.  

#### InSliceSli
	strarray.InSliceSli(l [][]string, s string, c int) bool  

Returns true if s is in column c in l  

### Set  
The set struct is a simple python-style set for strings.  

#### NewSet  
	NewSet()

Initializes new set.  

#### Len  
	set.Len()  

Returns length of set.  

#### Add  
	set.Add(value string)  

Adds string value to set.  

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

#### GetScanner  
	iotools.GetScanner(f *os.File) *bufio.Scanner  

Returns scanner for either gzipped or uncompressed file.  

#### Exists  
	iotools.Exists(path string) bool

Returns true if file or directory exists. Otherwise, returns false.  

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

#### GetDelim
	iotools.GetDelim(header string) string

Returns delimiter from header of a text file

#### WriteToCSV
	iotools.WriteToCSV(outfile, header string, results [][]string)  

Writes header and slice of string slices to comma seperated file.  
Prints error if line cannot be written but does not exit.  
