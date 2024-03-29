// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Short Discard
//
// 	1. Short declare two bool variables
//     (use multiple short declaration syntax)
//
//  2. Initialize both variables to true
//
//  3. Change your declaration and
//     discard the 2nd variable's value
//     using the blank-identifier
//
//  4. Print only the 1st variable
//
// EXPECTED OUTPUT
//  true
// ---------------------------------------------------------

func main() {
	// ADD YOUR DECLARATIONS HERE
	on, off := true, true

	_ = off

	// THEN UNCOMMENT THE CODE BELOW
	fmt.Println(on)
}
