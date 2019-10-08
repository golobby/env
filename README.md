[![GoDoc](https://godoc.org/github.com/golobby/env?status.svg)](https://godoc.org/github.com/golobby/env)
[![Build Status](https://travis-ci.org/golobby/env.svg?branch=master)](https://travis-ci.org/golobby/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/golobby/env)](https://goreportcard.com/report/github.com/golobby/env)
[![Coverage Status](https://coveralls.io/repos/github/golobby/env/badge.svg?branch=master)](https://coveralls.io/github/golobby/env?branch=master)

# Env
Env is a simple package to read and load environment variable files. 
It parses env files and returns their key/values in as a map. 
It also loads or overloads them into the operating system.

## Documentation

### Supported Versions
It requires Go `v1.11` or newer versions.

### Installation
To install this package run the following command in the root of your project

```bash
go get github.com/golobby/env
```

### Loading

Following example demonstrates how to load env file.

```go
vs, err := env.Load(".env")
fmt.Println(vs)
```

`vs` will be a `map[string]string]` of the environment variables. It also loads the variables to OS.

#### Example

Consider this env file:

```
APP_NAME=MyGoApp
APP_PORT=8585
```

To read and load it:

```go
vs, err := env.Load(".env")

fmt.Println(vs["APP_NAME"]) // MyGoApp
fmt.Println(vs["APP_PORT"]) // 8585
fmt.Println(os.Getenv("APP_NAME")) // MyGoApp
fmt.Println(os.Getenv("APP_PORT")) // 8585
```

### Overloading

The function `Load()` will load the variables to the operating system, it ignores the variables that have already 
existed in the operating system.

There is another function named `Overload()`. It's so similar to the `Load()` function but it overwrites the operating 
system variables.

```go
vs, err := env.Overload(".env")
```

## Contributors

* [@miladrahimi](https://github.com/miladrahimi)

## License

GoLobby Env is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
