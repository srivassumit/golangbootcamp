package main

// Exported name is the path (after $GOPATH) to the package
import "github.com/srivassumit/golangbootcamp/05/printer"

func main() {
	// This is not exported as the name doesn't start with uppercase letter.
	// printer.hello()
	// This is exported
	printer.Bye()
}
