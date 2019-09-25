## Section 5: Write a Library Package

Topics covered:

- Creating first library package
  - cannot be run
- Building the library package
  - `go build`
- installing the library package to `$GOPATH`
  - `go install` in the package directory.
  - This will install the library in path: `$GOPATH/pkg/<system_architecture>/github.com/<github_username>/project/package`
  - here the compiled `.a` file will be present. `.a` means compressed archive file.
    - This file depends on system (windows/mac etc.)
    - The compiled file name is same as the name of the library package
- exporting
  - export names when a package needs to be used in another package
  - to export a name just make its first letter an uppercase letter.
  - all files which import the package can use all names which start with uppercase letter.
  - `import` statement will automaticallly find the compiled package files which are imported.

**Note**: When you create an executable for project then use the name as `cmd`

### Exercise:

1. Create a new library
2. In it, create a function that returns Go version
3. Create a command and import your library
4. Call your function that returns Go version
5. Run your program
   Note: You need to run all the files.
   `go run main.go file2.go` and so on.

#### HINTS

Create your package function like this:

```
func Version() string {
    return runtime.Version()
}
```

#### EXPECTED OUTPUT

It should print the current Go version on your system.
