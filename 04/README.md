## Section 4: Learn Go Fundamentals: Statements and Expressions

Topics covered:

- Statement
  - Execution flow
  - Semicolons
    - Go implicitly puts semicolons behind the scene.
    - you can put multiple statements in one line if you manually put a semicolon in between.
- Expressions
  - code that computes one or multiple values.
  - in Go, an expression can return multiple values.
  - expression returns a value so it should be used in a statement to assign that value.
  - functions can also act like expressions eg `runtime.NumCPU()` function
- Go comments
  - not execute by compiler
  - single line : `//`
  - multi line: `/* .. */`
- Go Docs
  - `go doc` tool creates go docmentation automatically.

Exercises:

1. Shy Semicolons
   Observe how Go fixes semicolons for you.

2. Naked Expression
   Observe what happens when you use an expression without a statement.

3. Operators Combine
   Try using operators to combine expressions.

4. Print Go Version
   Use a package from Go Standard Library to print the current version of Go on your system. Or the system which will run your program.

5. Comment Out
   Learn how to comment out.

6. Use the GoDoc
   Try godoc yourself.
