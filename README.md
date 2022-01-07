# hello-go

hello-go

## go cmd

```sh
go version
go env
go clean
```

### start module

```sh
go mod init github.com/rosiirao/hello-go
go mod tidy # removes any dependencies the module might have accumulated that are no longer necessary.
go list -m all
go mod graph
go mod why -m
go get  # upgrade or downgrade to the correct version
```

## uninstall

Linux / macOS / FreeBSD

1. Delete the go directory.

    This is usually /usr/local/go.

2. Remove the Go bin directory from your PATH environment variable.

    Under Linux and FreeBSD, edit /etc/profile or $HOME/.profile. If you installed Go with the macOS package, remove the /etc/paths.d/go file.

## build

```sh
go build
```

## Run

> The go run command now allows a single import path, a directory name or a pattern matching a single package.
>
> This allows `go run pkg` or `go run dir`, most importantly `go run .`.

```sh
go run .
go run --work .
```

## Golang knowledge

### Re-declaration and reassignment

In a := declaration a variable v may appear even if it has already been declared, provided:

- this declaration is in the same scope as the existing declaration of v (if v is already declared in an outer scope, the declaration will create a new variable §[^1]),
- the corresponding value in the initialization is assignable to v, and
- there is at least one other variable that is created by the declaration.

[^1]: It's worth noting here that in Go the scope of function parameters and return values is the same as the function body, even though they appear lexically outside the braces that enclose the body.
