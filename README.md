# k8sjsonparser

This is a repo that i will use to make a basic parser function of k8s files from a folder.
It's a way for me to learn how to write go-lang and nothing else.

I will probably focus on creating something easy like a service and not much else with a very limited functionaillity.

In the long run i plan to createa a cli to call on this application as well but for now a go run will have to do the trick.

Is it worth setting up a seperate error channel https://www.atatus.com/blog/goroutines-error-handling/?

## Basic commands

I don't program enough so i forget basic commands.

```go run main.go```

```go test ./...```

## TODO

- [x] Move json parser to a seperate package
- [x] Use channels/go functions to parse json
- [x] Read all files from a folder
- [ ] Create cli
- [ ] Add tests
- [ ] Should the Structs be "global", find some best practice
