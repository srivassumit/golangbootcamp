// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"path"
)

// ---------------------------------------------------------
// EXERCISE: Discard The File
//
//  1. Print only the directory using `path.Split`
//
//  2. Discard the file part
//
// RESTRICTION
//  Use short declaration
//
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

func main() {
	// UNCOMMENT THE CODE BELOW:

	d, _ := path.Split("secret/file.txt")

	fmt.Println(d)
}
