package main

import "fmt"

func main() {
	speed := 100
	force := 2.5

	// This will give error as Force is float64 and speed is int
	// speed = speed * force

	// So, we can convert either force to int or speed to float64
	// e.g. -> speed = speed * int(force)
	// ^ this will result in loss of the fractional part

	// This will also not work as it does not convert type of speed variable to float64
	// So, we cannot assign a float64 to int.
	// speed = float64(speed) * force

	// we have to convert it back to an int to assign it to speed:
	speed = int(float64(speed) * force)

	// Also, this will also not work as you cannot multiple int to float64:
	//speed = int(speed * force)

	//Btw, This will work:
	force = force * 10
	// because here we are not multiplying it with an int, we are multiplying it with a constant
	// which does not have a type.

	fmt.Println(speed)
}
