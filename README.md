[![GoDoc](https://godoc.org/github.com/golobby/env/v2?status.svg)](https://godoc.org/github.com/golobby/env/v2)
[![CI](https://github.com/golobby/env/actions/workflows/ci.yml/badge.svg)](https://github.com/golobby/env/actions/workflows/ci.yml)
[![CodeQL](https://github.com/golobby/env/workflows/CodeQL/badge.svg)](https://github.com/golobby/env/actions?query=workflow%3ACodeQL)
[![Go Report Card](https://goreportcard.com/badge/github.com/golobby/env)](https://goreportcard.com/report/github.com/golobby/env)
[![Coverage Status](https://coveralls.io/repos/github/golobby/env/badge.svg?branch=master)](https://coveralls.io/github/golobby/env?branch=master)

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
    Debug bool      `env:"DEBUG"`
    App struct {
        Name string `env:"APP_NAME"`
        Port int16  `env:"APP_PORT"`
    }
    Database struct {
        Name string `env:"DB_NAME"`
        Port int16  `env:"DB_PORT"`
        User string `env:"DB_USER"`
        Pass string `env:"DB_PASS"`
    }
}

c := Config{}

err := env.Load(&c)
if err != nil {
    panic(err)
}
```

## See Also
* ***[golobby/dotenv](https://github.com/golobby/dotenv)**: A lightweight package for loading dot env (.env) files into structs in Go projects*

## License
GoLobby Env is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
