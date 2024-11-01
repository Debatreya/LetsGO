# Build for Windows, Linux, Mac
Create your stand alone program.

## Steps:
- Use Command `go env`
- Search your `GOOS` variable
- For building use `go build` 
    It will take the default `GOOS`. If you want to build for mac `GOOS="darwin" go build`

Lets create a standalone program that will give us the current DATE-TIME-DAY whenever we run it
- Initialize and Follow [main.go](./main.go). 
- Run: `go build -o myTime main.go`
- Run: `myTime`. If the file lack extension rename it with  `myTime.exe`


---
[Back to Contents](../../Readme.md)
[Next Tutorial](../09tut/index.md)