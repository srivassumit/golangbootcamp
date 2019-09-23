## Section 3: Learn Go Fundamentals - Packages, Scopes and importing

Topics covered:

- Packages:
  - all files of same package must be in same folder.
  - all files in same folder should be in same pacakge (convention)
  - to make a package executable, it should be called `main`.
  - package line should be first line of code.
  - in one file, package can come only once.
- Compiling and running go programs with multiple files:
  - `go run *.go` will compile all files and run all files which contain main method
  - `go run` can also take multiple files like this: `go run main.go hey.go bye.go` and it will execute all files which have `func main()` in them.
- Executable vs Library packages.
  - executable package must always be called main and it should also contain (only one) `func main()`.
  - Library package can be called anything and there is no `func main()` in library package.
  - library package cannot be executed but can only be imported (e.g. `fmt`).
  - executable package can only be run, but cannot be imported.
- Scopes:
  - file scope - visible to file only
    - eg `import "fmt"` - visible only to the current file.
  - pacakge scope - visible to entire package..
    - eg global declarations in a file.
    - if we are using method from another file in same package, the the entire package needs to be run using `go run *.go`
  - block scope - anything inside a block i.e. within `{ ... }` other pacakges and other blocks can't see this.
- Importing:
  - import can only be done at a file level not at package level.
  - ot use same import again in a different file in same package, it needs to be imported again.
- Rename imported packages
  - same package cannot be imported twice with the same name. Each import requires a unique name.
  - Declare a new name for the imported package as: `import f "fmt"`
  - This syntax is used when we need to import two package which have the same name
