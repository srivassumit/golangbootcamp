package main

import "fmt"

func main() {
	// Declare and initialize multiple variables using multiple short declaration
	safe, speed := true, 50
	fmt.Println(safe, speed)

	name, lastName := "Nokila", "Tesla"
	fmt.Println(name, lastName)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	// Declaring a lot of variables on same line can become unreadable.
	degree, ratio, heat := 10.55, 30.5, 20.
	fmt.Println(degree, ratio, heat)

	nFiles, valid, msg := 10, true, "hello"
	fmt.Println(nFiles, valid, msg)

}
