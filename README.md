# dotenv

The `github.com/talut/dotenv` package is a Go library for managing environment variables. It provides a set of functions to retrieve environment variables with different types, including string, boolean, integer, float, and duration. It also includes "Must" versions of these functions that panic if the environment variable is not set or cannot be parsed correctly.

## Functions

- `GetString(key, fallback string) string`: Retrieves the value of the environment variable named by the key. If the value is not set, the fallback value is returned.
- `GetBool(key string, fallback bool) bool`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as a boolean, the fallback value is returned.
- `GetInt(key string, fallback int) int`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as an integer, the fallback value is returned.
- `GetFloat(key string, fallback float64) float64`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as a float, the fallback value is returned.
- `GetDuration(key string, fallback time.Duration) time.Duration`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as a duration, the fallback value is returned.
- `MustGetString(key string) string`: Retrieves the value of the environment variable named by the key. If the value is not set, a panic occurs.
- `MustGetBool(key string) bool`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as a boolean, a panic occurs.
- `MustGetInt(key string) int`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as an integer, a panic occurs.
- `MustGetFloat(key string) float64`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as a float, a panic occurs.
- `MustGetDuration(key string) time.Duration`: Retrieves the value of the environment variable named by the key. If the value is not set or cannot be parsed as a duration, a panic occurs.

## Usage

Import the `dotenv` package into your Go code using the following import statement:

```go
 import "github.com/talut/dotenv"
```
