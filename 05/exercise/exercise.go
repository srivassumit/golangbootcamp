package exercise

import (
	"fmt"
	"runtime"
)

// Version exported function
func Version() {
	fmt.Println(runtime.Version())
}
