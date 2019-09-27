// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool literals
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

func main() {
	// Use fmt.Println()

	fmt.Println(1, -234, 64, 42, -76)
	fmt.Println(1.0, .5-12., 133.42, -98.12, .8)
	fmt.Println(true, false)
	fmt.Println("Sumit")
	fmt.Println("Non English characters &µαξΨ")
}
