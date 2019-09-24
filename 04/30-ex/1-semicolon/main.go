// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"
// ---------------------------------------------------------
// EXERCISE: Shy Semicolons
//
//  1. Try to type your statements by separating them using
//     semicolons
//
//  2. Observe how Go fixes them for you
//
// ---------------------------------------------------------

func main() {
  // Below line is automatically split to 2 separate lines and semicolon is removed when saved.
  fmt.Println("a"); fmt.Println("b")
  // To avoid auto format use Ctrl-K + Ctrl-Shift-S in VSCode
}
