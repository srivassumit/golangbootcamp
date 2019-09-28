package main

import "fmt"

func main() {
	// Already Declared variable
	var safe bool

	// Redeclaring using short declaration
	// Here there must be atleast one other new variable or it won't work.
	safe, speed := true, 50

	// So, doing this: won't work
	// It gives error `No new variables on left side of :=`
	// safe := false

	fmt.Println(speed, safe)

	name := "Nokila"
	fmt.Println(name)

	// Redeclaration
	name, age := "Marie", 66
	fmt.Println(name, age)

	// Value assignment and declaring a new variable
	name = "Albert"
	birth := 1879
	fmt.Println(name, birth)

	// This is same as above, it redeclares name and declares a new variable birth.
	// name, birth := "Albert", 1879
	// So, redeclaration is same as an assignment and a declaration in the same statement.

	// You cannot change variable type while redeclaration though
	// So, this won't work:
	// name, death := 16, 1990
}
