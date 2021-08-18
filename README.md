[![GoDoc](https://godoc.org/github.com/golobby/env?status.svg)](https://godoc.org/github.com/golobby/env)
[![Build Status](https://travis-ci.org/golobby/env.svg?branch=master)](https://travis-ci.org/golobby/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/golobby/env)](https://goreportcard.com/report/github.com/golobby/env)
[![Coverage Status](https://coveralls.io/repos/github/golobby/env/badge.svg?branch=master)](https://coveralls.io/github/golobby/env?branch=master)

# Env
GoLobby Env is a simple package to read environment variable files and load them into the OS variables. 
It parses env files and returns their key/values as a `map[string]string`. 
It also loads or overloads them into the OS variables.

## Documentation

### Supported Versions
It requires Go `v1.11` or newer versions.

### Installation
To install this package run the following command in the root of your project

```bash
go get github.com/golobby/env
```

### Loading

Following example demonstrates how to read and load a env file.

```go
vs, err := env.Load(".env")
fmt.Println(vs)
```

The filename you pass to `Load()` function must be env file relative path like `.env` or `path/to/.env`.
`vs` variable in the example will be a `map[string]string]` of the environment variables. 
It also loads the variables to OS variables.

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

The function `Load()` will load the variables into the operating system, it ignores the variables that have already 
existed in the OS variables.

There is another function named `Overload()`. It's so similar to the `Load()` function but the difference is that 
it overwrites the OS variables.

The example below shows how it change the OS variables.

```go
err := os.Setenv("APP_NAME", "TheOldName")
fmt.Println(os.Getenv("APP_NAME")) // TheOldName

vs, err := env.Overload(".env")
fmt.Println(os.Getenv("APP_NAME")) // MyGoApp
```

### Syntax
This package works with standard syntax of environment variable as explained below.

Example of env file:

```
# This is comment

APP_NAME=MyGoAPP
APP_PORT = 8585
```

Rules:
* Lines that start with `#` will be ignored.
* Space around keys and values will be trimmed.
* Empty lines will be ignored.

## Contributors

* [@miladrahimi](https://github.com/miladrahimi)

## License

GoLobby Env is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
