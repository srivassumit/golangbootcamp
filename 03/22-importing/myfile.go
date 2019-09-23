package main

// In a new file in the same pacakge (or any other package),
// we'll need to import fmt again if we need to use it.
import (
	"fmt"

	// import the same package again as `f`
	f "fmt"
)

func main() {
	fmt.Println("Hello!")
	f.Println("Bye!")
}
