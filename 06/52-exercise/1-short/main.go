// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Short Declare
//
//  Declare and then print four variables using
//  the short declaration statement.
//
// EXPECTED OUTPUT
//  i: 314 f: 3.14 s: Hello b: true
// ---------------------------------------------------------

func main() {
	// ADD YOUR DECLARATIONS HERE
	i, f, s, b := 314, 3.14, "Hello", true

	// THEN UNCOMMENT THE CODE BELOW
	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
}
