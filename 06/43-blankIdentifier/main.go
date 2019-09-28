package main

// This will not give unused variable error
// even though it is not used in this file,
// because it may still be getting used
// from some other file.
var c int

func main() {
	// This will give unused variable error at compile time
	// var a int

	// This is also unused, but
	var b int
	// using the blank identifier we can avoid the unused variable error.
	_ = b
}
