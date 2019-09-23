## Section 2: * PART 1: Write your first Go Program

Topics covered:

- Writing first go code
- The `$GOPATH` environment variable:
  - Linux: `/usr/<username>/go`
  - Mac: `/usr/<username>/go`
  - Windows: `C:\Users\<username>\go`
- Folder structure workspace and placing Go files:
  - `$GOPATH/src/*`
  - Recommended: `$GOPATH/src/github.com/<githubusername>/<projectname>/`
- Compiling with `go build <filename>`
- Check where file can be executed by using `file <executable_file_name>`
- Running code with `go run <filename>`
  - check details by using `go run -x <filename>`
- Check the Go Environment by typing `go env`
- Cross platform compiling (from linux/unix terminal):

  > Create an OS X executable: `GOOS=darwin GOARCH=386 go build <filename>`
  > Create a Windows executable: `GOOS=windows GOARCH=386 go build <filename>`
  > Create a Linux executable: `GOOS=linux GOARCH=arm GOARM=7 go build <filename>`

  > You can find the full list of environments in [here](https://golang.org/doc/install/source#environment):

  > NOTE: If you're using the command prompt or the powershell, you may need to use it like this: `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

- Go online [documentation](https://golang.org/pkg)
- [Golang tour](https://tour.golang.org/)
