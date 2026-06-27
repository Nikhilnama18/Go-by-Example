# Go-by-Example
Hands code for Go-by-Example

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
