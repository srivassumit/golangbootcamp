package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		speed int
		now   time.Time
	)

	// Works similar to how short declaration works
	// not recommended to assign more than 3 in a single statement
	speed, now = 100, time.Now()

	fmt.Println(speed, now)

	// Variables can be easily swapped:
	prevSpeed := 50

	speed, prevSpeed = prevSpeed, speed

	fmt.Println(speed, prevSpeed)
}
