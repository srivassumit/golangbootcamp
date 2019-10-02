package main

import "fmt"

func main() {
	var apple int
	var orange int32

	// int cannot be assigned to int32 or vice versa even though they are both int.
	// apple = orange
	// orange = apple

	apple = int(orange)
	orange = int32(orange)

	// but uncompatible types cannot be converted. So, this won't work
	// isDelicious := bool(orange)

	orange = 65

	// However, this will work as it will just assign the ASCII value corresponding to that number to the String variable
	color := string(orange)
	// 65 = A so color waill have the value A

	// This will work too and print the symbol corresponding to ASCII 165
	color2 := string(165)

	// This will also work:
	str3 := string([]byte{104, 105})
	// It will create a new string with characters corresponding to the string values of each byte.
	// 104 = `h` and 105 = `i` So, str3 will contain `hi`

	fmt.Println(apple, orange, color, color2, str3)
}
