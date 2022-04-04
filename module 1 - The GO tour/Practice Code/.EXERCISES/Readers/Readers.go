package main

//This exercise is executable in go background uncommenting
//the package imported, the return of the error when the array is
//empty and the execution of the Validate function in func main
//import (	"golang.org/x/tour/reader")

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(a []byte) (int, error) {
	a[0] = 'A' //returning value 1 and nil error
	return 1, nil
}

func main() {
	//method that verifies the requirement of the infinite loop
	//reader.Validate(MyReader{})
}
