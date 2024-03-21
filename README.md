# dotenv

[![Go Reference](https://pkg.go.dev/badge/github.com/talut/dotenv.svg)](https://pkg.go.dev/github.com/talut/dotenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/talut/dotenv)](https://goreportcard.com/report/github.com/talut/dotenv)
[![License](https://img.shields.io/github/license/talut/dotenv)](LICENSE)

The `github.com/talut/dotenv` package is a Go library for managing environment variables. It provides a set of functions
to retrieve environment variables with different types, including string, boolean, integer, float, and duration. It also
includes "Must" versions of these functions that panic if the environment variable is not set or cannot be parsed
correctly.

## Functions

- `ClearCache()`: Clears the cache of environment variables.
- `GetString(key, fallback string) string`: Retrieves the value of the environment variable named by the key. If the
  value is not set, the fallback value is returned.
- `GetBool(key string, fallback bool) bool`: Retrieves the value of the environment variable named by the key. If the
  value is not set or cannot be parsed as a boolean, the fallback value is returned.
- `GetInt(key string, fallback int) int`: Retrieves the value of the environment variable named by the key. If the value
  is not set or cannot be parsed as an integer, the fallback value is returned.
- `GetFloat(key string, fallback float64) float64`: Retrieves the value of the environment variable named by the key. If
  the value is not set or cannot be parsed as a float, the fallback value is returned.
- `GetDuration(key string, fallback time.Duration) time.Duration`: Retrieves the value of the environment variable named
  by the key. If the value is not set or cannot be parsed as a duration, the fallback value is returned.
- `MustGetString(key string) string`: Retrieves the value of the environment variable named by the key. If the value is
  not set, a panic occurs.
- `MustGetBool(key string) bool`: Retrieves the value of the environment variable named by the key. If the value is not
  set or cannot be parsed as a boolean, a panic occurs.
- `MustGetInt(key string) int`: Retrieves the value of the environment variable named by the key. If the value is not
  set or cannot be parsed as an integer, a panic occurs.
- `MustGetFloat(key string) float64`: Retrieves the value of the environment variable named by the key. If the value is
  not set or cannot be parsed as a float, a panic occurs.
- `MustGetDuration(key string) time.Duration`: Retrieves the value of the environment variable named by the key. If the
  value is not set or cannot be parsed as a duration, a panic occurs.

## Installation

```shell
go get github.com/talut/dotenv
```

## Usage

Import the `import "github.com/talut/dotenv"` package into your Go code using the following import statement:

Mote: The second parameters in the Get functions are fallback values that will be used
if the environment variable is not set. The MustGet functions will panic if the environment variable is not set or
cannot be parsed correctly.

```go
package main

import (
	"fmt"
	"time"

	"github.com/talut/dotenv"
)

func main() {
	// Example usage of GetString
	fmt.Println(dotenv.GetString("STRING_ENV_VAR", "default"))

	// Example usage of GetBool
	fmt.Println(dotenv.GetBool("BOOL_ENV_VAR", false))

	// Example usage of GetInt
	fmt.Println(dotenv.GetInt("INT_ENV_VAR", 0))

	// Example usage of GetFloat
	fmt.Println(dotenv.GetFloat("FLOAT_ENV_VAR", 0.0))

	// Example usage of GetDuration
	fmt.Println(dotenv.GetDuration("DURATION_ENV_VAR", 1*time.Second))

	// Example usage of MustGetString
	fmt.Println(dotenv.MustGetString("MUST_STRING_ENV_VAR"))

	// Example usage of MustGetBool
	fmt.Println(dotenv.MustGetBool("MUST_BOOL_ENV_VAR"))

	// Example usage of MustGetInt
	fmt.Println(dotenv.MustGetInt("MUST_INT_ENV_VAR"))

	// Example usage of MustGetFloat
	fmt.Println(dotenv.MustGetFloat("MUST_FLOAT_ENV_VAR"))

	// Example usage of MustGetDuration
	fmt.Println(dotenv.MustGetDuration("MUST_DURATION_ENV_VAR"))

	// Example usage of ClearCache
	dotenv.ClearCache()
}
```

Please replace `"STRING_ENV_VAR", "BOOL_ENV_VAR", "INT_ENV_VAR", "FLOAT_ENV_VAR", "DURATION_ENV_VAR", "
MUST_STRING_ENV_VAR", "MUST_BOOL_ENV_VAR", "MUST_INT_ENV_VAR", "MUST_FLOAT_ENV_VAR", and "MUST_DURATION_ENV_VAR"` with
your actual environment variable names.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
