# go-tools  

## Commonly used GO functions (intended for personal use, but feel free to use)  

### To download and install to gopath:  
	go get github.com/icwells/go-tools/iotools  
	go get github.com/icwells/go-tools/strarray  

	go install github.com/icwells/go-tools/iotools  
	go install github.com/icwells/go-tools/strarray  

### strarray contains functions for working with slices and maps of strings  

#### InSliceStr
InSliceStr(l []string, s string) bool  

Returns true if s is in l.  

#### InSliceSli
InSliceSli(l [][]string, s string, c int) bool  

Returns true if s is in column c in l  

#### InMapStr  
InMapStr(m map[string]string, s string) bool  

Returns true if s is in m keys. 

#### InMapSli  
InMapSli(m map[string][]string, s string) bool  

Returns true if s is in m keys. 


#### InMapMapStr  
InMapMapStr(m map[string]map[string]string, s string) bool  

Returns true if s is a key in outer map.  

#### InMapMapSli
InMapMapSli(m map[string]map[string][]string, s string) bool  

Returns true if s is a key in outer map.  

### iotools wraps open and create file functions with error handling  

#### OpenFile
OpenFile(file string) *os.File  

Returns pointer to File type. Prints and error and exits if file cannot be opened.  

#### CreateFile
CreateFile(file string) *os.File   

Creates new file and returns pointer. Prints and error and exits if file cannot be created.

#### WriteToCSV
WriteToCSV(outfile, header string, results [][]string)  

Writes header and slice of string slices to comma seperated file.  
Prints error if line cannot be written but does not exit.  
