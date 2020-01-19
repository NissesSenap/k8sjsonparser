# k8sjsonparser

This is a repo that i will use to make a basic parser function of k8s files from a folder.
It's a way for me to learn how to write go-lang and nothing else.

I will probably focus on creating something easy like a service and not much else with a very limited functionaillity.

In the long run i plan to createa a cli to call on this application as well but for now a go run will have to do the trick.

## Basic commands

I don't program enough so i forget basic commands.

### Start application

```go run main.go```

### Run tests

```go test ./...```

### Vet deeper test of "syntax"

Vet examines Go source code and reports suspicious constructs
Allways good to run on all code bases

```go vet```

### Run tests + coverage

```go test ./... -cover```

### Visualize test coverage html

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

### Find potential memory leaks

-race tests to see if memory is used at the same time.

```go test -race ./...```

## Tips

### Add parallel testing

If your tests can handle it add bellow to get
your application to run on all your cores

```t.parallel()```

## Questions for my self

Is it worth setting up a seperate [error channel](https://www.atatus.com/blog/goroutines-error-handling/)?

Talking to an old colleague it's better to use [error groups](https://godoc.org/golang.org/x/sync/errgroup)

## TODO

- [x] Move json parser to a seperate package
- [x] Use channels/go functions to parse json
- [x] Read all files from a folder
- [ ] Create cli
- [ ] Add tests
- [ ] Should the Structs be "global", find some best practice
