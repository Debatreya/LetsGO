# What is 'MOD' tool in GoLANG?

[Go Module Blog](https://go.dev/blog/using-go-modules)
Initialize and Follow [main.go](./main.go)

- go mod tidy (tidy up the dependencies)
- go mod verify (checks hashes in go.sum)
- go list (lists dependencies)
- go list all (lists all downloaded dependencies)
- go list -m all (lists responsible dependencies)
- go list -m -versions github.com/gorilla/mux
- go mod vendor (creates /vendor which is something like node_modules, 
to use it `go run -mod=vendor main.go`)

---
[Back to Contents](../../Readme.md)
[Next Tutorial](../30tut/index.md)