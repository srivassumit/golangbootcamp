package main

import (
	"fmt"
	"runtime"
)

func main() {
	// "Hello!" below is an expression
	fmt.Println("Hello!")

	// "Hello!" + "!" below is an expression and both "Hello!" and "!" are expressions too
	fmt.Println("Hello!" + "!")

	// Println is a statement and runtime.NumCPU() is an expression
	fmt.Println(runtime.NumCPU())
}
