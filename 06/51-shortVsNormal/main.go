package main

import "fmt"

// At package scope you can only use normal declaration
// as short declaration does not work here.
//version := 0
var version int

func main() {
	// score := 0 // Don't do this if you don't know initial value
	var score int // Do this instead as 0 will already be assigned as zero value

	// Use normal declaration for grouping variables:
	var (
		// grouped declarations for related variables
		video string

		// closeley related
		duration int
		current  int
	)

	// If you know initial value, then use short declaration
	// var width, height = 100, 50 // Don't do this if you know initial values
	width, height := 100, 50 //Use this instead

	// use short declaration for redeclaration
	// width = 50 // Don't do this.
	// color := "green"
	width, color := 150, "green" // Do this instead as it is more concise

	fmt.Println(score, video, duration, current, width, height, color)
}
