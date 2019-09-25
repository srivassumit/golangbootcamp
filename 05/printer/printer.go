package printer

import "fmt"

// this won't be exported
func hello() {
	fmt.Println("Hello!")
}

// Bye this will be exported
func Bye() {
	fmt.Println("Bye!")
}
