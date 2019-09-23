package main

// File scope visible to this file only
import "fmt"

// Package scope visible to all files in current package (main)
const ok = true

// Package scope
func nope() {
	// Block scope of block nope. Not visible outside this scope.
	const Nok = false
	//
	var yellow = "Yellow!"
	// Blank identified to avoid the error on unused variable above.
	_ = yellow
}

// Cannot declare here again as it is declared already in bye.go
// func bye() {
// 	fmt.Println("Bye 2!")
// }

// Package scope.
func main() {
	// Block scope. only visible within the block defined within { .. }
	var hello = "Hello"
	fmt.Println(hello, ok)

	// Not visible - yellow and nok because they are in the nope scope
	// 	fmt.Println(yellow, nok)

	// bye() from bye.go can be called here as it is in package scope
	bye()
}
