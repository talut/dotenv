package dotenv

import (
	"log"
	"os"
	"strconv"
	"time"
)

// cache is a map that stores the values of environment variables.
// The key is the name of the environment variable and the value is the value of the environment variable.
// This is used to avoid repeated lookups of environment variables, which can be expensive.
var cache = make(map[string]string)

// GetString retrieves the value of the environment variable named by the key.
// If the value is not set, the fallback value is returned.
// The value is cached to avoid repeated lookups.
func GetString(key, fallback string) string {
	value, exists := cache[key]
	if !exists {
		value, exists = os.LookupEnv(key)
		if !exists || value == "" {
			value = fallback
		}
		cache[key] = value
	}
	return value
}

// GetBool retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as a boolean, the fallback value is returned.
// The value is cached to avoid repeated lookups.
// If the value is set but cannot be parsed as a boolean, a warning logged.
func GetBool(key string, fallback bool) bool {
	value, exists := cache[key]
	if !exists {
		value, exists = os.LookupEnv(key)
		if !exists {
			return fallback
		}
		cache[key] = value
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Printf("Failed to parse %s as bool: %v", key, err)
		return fallback
	}
	return boolValue
}

// GetInt retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as an integer, the fallback value is returned.
// The value is cached to avoid repeated lookups.
// If the value is set but cannot be parsed as an integer, a warning logged.
func GetInt(key string, fallback int) int {
	value, exists := cache[key]
	if !exists {
		value, exists = os.LookupEnv(key)
		if !exists {
			return fallback
		}
		cache[key] = value
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Failed to parse %s as int: %v", key, err)
		return fallback
	}
	return intValue
}

// GetFloat retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as a float, the fallback value is returned.
// The value is cached to avoid repeated lookups.
// If the value is set but cannot be parsed as a float, a warning logged.
func GetFloat(key string, fallback float64) float64 {
	value, exists := cache[key]
	if !exists {
		value, exists = os.LookupEnv(key)
		if !exists {
			return fallback
		}
		cache[key] = value
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Failed to parse %s as float: %v", key, err)
		return fallback
	}
	return floatValue
}

// GetDuration retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as a duration, the fallback value is returned.
// The value is cached to avoid repeated lookups.
// If the value is set but cannot be parsed as a duration, a warning logged.
func GetDuration(key string, fallback time.Duration) time.Duration {
	value, exists := cache[key]
	if !exists {
		value, exists = os.LookupEnv(key)
		if !exists {
			return fallback
		}
		cache[key] = value
	}
	durationValue, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Failed to parse %s as duration: %v", key, err)
		return fallback
	}
	return durationValue
}

// MustGetString retrieves the value of the environment variable named by the key.
// If the value is not set, a panic occurs.
func MustGetString(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " is not set")
	}
	return value
}

// MustGetBool retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as a boolean, a panic occurs.
func MustGetBool(key string) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " is not set")
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic("Failed to parse " + key + " as bool: " + err.Error())
	}
	return boolValue
}

// MustGetInt retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as an integer, a panic occurs.
func MustGetInt(key string) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " is not set")
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic("Failed to parse " + key + " as int: " + err.Error())
	}
	return intValue
}

// MustGetFloat retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as a float, a panic occurs.
func MustGetFloat(key string) float64 {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " is not set")
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic("Failed to parse " + key + " as float: " + err.Error())
	}
	return floatValue
}

// MustGetDuration retrieves the value of the environment variable named by the key.
// If the value is not set or cannot be parsed as a duration, a panic occurs.
func MustGetDuration(key string) time.Duration {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " is not set")
	}
	durationValue, err := time.ParseDuration(value)
	if err != nil {
		panic("Failed to parse " + key + " as duration: " + err.Error())
	}
	return durationValue
}
