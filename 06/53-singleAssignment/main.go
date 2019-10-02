package main

import "fmt"

func main() {
	var speed int
	fmt.Println(speed)

	// Go is strongly types - you can't change data type
	// speed = "100"

	speed = 100
	fmt.Println(speed)

	speed = speed - 25
	fmt.Println(speed)

	// This won't work either as Go is strongly types so 1 and true are not same.
	// var running bool
	// running = 1

	// This is also not allowed even between numerical types
	var force float64
	// speed = force

	// But, this is allowed
	// because a numberic constant can be converted to a numeric type automatically,
	// Here, 1 is a constant and has no type..
	// So, Go is not converting from int to float64, but from a numerical constant to float64
	force = 1
	//
	fmt.Println(force)
	var (
		name   string
		age    int
		famous bool
	)

	name = "John"
	age = 32
	famous = false

	fmt.Println(name, age, famous)

	name = "Somebody"
	age = 22
	famous = true

	fmt.Println(name, age, famous)

	var prevName string
	prevName = name // no by ref here, value is copied to new variable

	name = "Einstein"

	fmt.Println("prev Name:", prevName)
	fmt.Println("new Name:", name)
}
