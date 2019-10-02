package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	// Getting back two values from the function
	dir, file = path.Split("css/main.css")

	// Getting back two values but ignoring one value
	// _, file = path.Split("css/main.css")

	// We could have directly used a short assignment here
	// _, file := path.Split("css/main.css")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}
