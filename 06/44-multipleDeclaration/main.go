package main

import "fmt"

func main() {
	// Declaring multiple variables of the same type
	// Parellel multiple declaration
	var a, b, c int

	// Declaring multiple variables of different type
	var (
		x int
		y float64
		//related variables may be grouped together
		z      bool
		zz, tt string
	)

	fmt.Println(a, b, c, x, y, z, zz, tt)
}
