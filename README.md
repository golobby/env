[![GoDoc](https://godoc.org/github.com/golobby/env?status.svg)](https://godoc.org/github.com/golobby/env)
[![Build Status](https://travis-ci.org/golobby/env.svg?branch=master)](https://travis-ci.org/golobby/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/golobby/env)](https://goreportcard.com/report/github.com/golobby/env)
[![Coverage Status](https://coveralls.io/repos/github/golobby/env/badge.svg?branch=master)](https://coveralls.io/github/golobby/env?branch=master)
[![CodeQL](https://github.com/golobby/cast/workflows/CodeQL/badge.svg)](https://github.com/golobby/cast/actions?query=workflow%3ACodeQL)

# Env

GoLobby Env is a lightweight package for loading OS environment variables into structs in Go projects.

## Documentation
### Supported Versions
It requires Go `v1.11` or newer versions.

### Installation
To install this package run the following command in the root of your project
```bash
go get github.com/golobby/env/v2
```

### Usage Example
The following example demonstrates how to use GoLobby Env package.

```go
type Config struct {
    AppName string  `env:"APP_NAME"`
    AppPort int64   `env:"APP_PORT"`
    IsAdmin bool    `env:"IS_ADMIN"`
}

c := &Config{}

err := env.Load(&c)
if err != nil {
    panic(err)
}
```

## License
GoLobby Env is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
