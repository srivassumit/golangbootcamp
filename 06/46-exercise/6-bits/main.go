// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// 3. After you've done, check out the solution
//    and read the comments there
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

func main() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	var b bool
	var s string
	var r rune  // also a numeric type
	var bt byte // also a numeric type
	fmt.Println(i, i8, i16, i32, i64, f32, f64, c64, c128, b, s, r, bt)
}
