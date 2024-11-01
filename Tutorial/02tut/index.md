# Steps for using `Go` commands:
- Open Integrated Terminal
- Write 
    ```
    go
    ```
    To get a list of commands that you can use.
- You will get a `go help` and `go help <topic>` command which will help you with further details about using each command.

Eg
## Steps for using `GOPATH` command:
- Write `go help gopath`.
- Read the details. And lets use it.
- `go env GOPATH` this **GOPATH** is a case sensitive variable, not a command.
- Change Directory into the path given by the above command. And explore the different directories, modules, go has.

# Whats in that 'Go' folder?

/
    bin
        gopls.exe (for windows)
    pkg
        mod
            cache/
            github.com
            goland.org
            honnef.co
            mvdan.cc
        sumdb
            sum.golang.org


---
[Back to Contents](../../Readme.md)
[Next Tutorial](../03lexer/index.md)