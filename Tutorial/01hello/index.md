## Steps
- Go to Integrated Terminal in the same folder.
- Use Command to download the modules 
    ```
    go mod init hello
    ```
- Create [`main.go`](./main.go)
- Follow the code.
- Run the code with 
    ```
    go run main.go
    ```

### What is `go mod init hello`?
    The command `go mod init hello` initializes a new Go module in the current directory module name is `hello`. It creates a go.mod file that tracks the module's dependencies. This is similar to running `npm init` in a Node.js project to create a package.json file.

### What is `go run main.go`?
    It runs the `main.go` file. It just runs and doest not create any executables.

---
[Back to Contents](../../Readme.md)
[Next Tutorial](../02tut/index.md)