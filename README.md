# Go-by-Example
Hands on code for Go-by-Example

### package
* Package is a way to organise/group files in go
* Files within a package are complied together

### Visibility
* Go doesn't have private/public access specifiers
* Capitalisation determines the visibility of the method
* A method starting with Captial letter is exported, where as a method strating with small case is package private.
* Package private means, the method is accessible with the same package but not outside of it.

### Execution
* Go has 2 phases Compile & Execution
* Go code is first compiled & then executed
* Go's execution starts from main method present inside package main.
* To complie & execute together we use `go run . or (file-path)`
* For compile only we use `go build . or (file-path)`

### Data types
* Go offers various types of data types
    * Integer: int, int8, int16, int32, int64
    * UnSigned: uint, uint8, uint16, uint32, uint64, uintptr
    * Float: float32, float64
    * Complex: complex64, complex128
    * Boolean: bool
    * Text: string
    * Collection: array, slice, map
    * Custom data type: struct
    * Address: pointer
    * Behaviour: interface
    * Concurrency: chan i.e (channel)
    * Additional: byte, rune
* Go has first class function support that means, functions can be stored as variables & passed.

### Pass by Value
* Go is techinically pass by value.
* Values are copied & passed onto other variables
* But for some data types the values are references. So the address is copied changing value will effect everywhere.

* Value type are: int, float, bool, array, struct
* Reference type are: slice, map, func, pointer, interface
