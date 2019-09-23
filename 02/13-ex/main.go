package main

import "fmt"

func main() {

	fmt.Println("My name: Sumit S")
	fmt.Println("My friend's name: Friend")

	fmt.Print("Testing Print")
	fmt.Println("Testing Println")

	fmt.Print("Multiple", "values", "using", "Print")
	fmt.Println("Multiple", "values", "using", "Println")

	// Gives error
	// fmt.Println(removing double quotes)

	// Gives Error
	// fmt.Println('single Quotes don\'t work either')

}

// Importing here gives error.
// import "fmt"

// Needs to be the first line
// package main
