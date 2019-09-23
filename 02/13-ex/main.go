package main

import "fmt"

func main() {

	// You can check the Go Environment by typing `go env` on terminal

	// You can create executables for different platform as follows:

	// Create an OS X executable (from linux/unix terminal):
	// GOOS=darwin GOARCH=386 go build

	// Create a Windows executable (from linux/unix terminal):
	// GOOS=windows GOARCH=386 go build

	// Create a Linux executable (from linux/unix terminal):
	// GOOS=linux GOARCH=arm GOARM=7 go build

	// You can find the full list in here:
	// https://golang.org/doc/install/source#environment

	// 😱NOTE: If you're using the command prompt or the powershell, you may need to use it like this:
	// cmd /c "set GOOS=darwin GOARCH=386 && go build"

	fmt.Println("My name: Sumit S")
	fmt.Println("My friend's name: Friend")

	fmt.Print("Testing Print")
	fmt.Println("Testing Println")

	fmt.Print("Multiple", "values", "using", "Print")
	fmt.Println("Multiple", "values", "using", "Println")

	// Gives error
	// fmt.Println(removing double quotes)

	// Gives Error
	// fmt.Println('single Quotes don\'t work either')
}

// Importing here gives error.
// import "fmt"

// Needs to be the first line
// package main
