package main

import "fmt"

func main() {
	// if typing like this on a new line, an extra comma is needed after the last value
	fmt.Println(
		-200, -10, 0, 50, 100,
	)
	// if typing like below, the extra comma is not needed
	fmt.Println(-200, -10, 0, 50, 100)

	// for floats the (non-significant) zero does not need to be written
	fmt.Println(-50.5, -20, 20., .9, -0., 100.5)
	// While printing the above output, Go will omitthe decimal in 20. and -0. and add a zero before .9
	// how it print can be changed using printf instead of Println.

	fmt.Println(true, false)

	fmt.Println(
		"", // empty string - nothing is printed for this on the console
		"hi",
		"✔✔✔", // non english characters
	)
}
