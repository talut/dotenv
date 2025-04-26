package dotenv

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// cache is a map that stores the values of environment variables.
// The key is the name of the environment variable and the value is the value of the environment variable.
// This is used to avoid repeated lookups of environment variables, which can be expensive.
var cache = make(map[string]string)

// ClearCache clears the cache of environment variables.
func ClearCache() {
	cache = make(map[string]string)
}

// Load loads environment variables from one or more .env files.
// Files are loaded in the order provided. If a key exists in multiple files,
// the value from the last file will be used.
// If no filenames are provided, it attempts to load from the default ".env" file.
// Files that don't exist are skipped without error.
func Load(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}

	for _, filename := range filenames {
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			continue
		}
		contents, err := os.ReadFile(filename)
		if err != nil {
			return err
		}

		lines := strings.Split(string(contents), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			idx := strings.Index(line, "=")
			if idx == -1 {
				continue // Skip lines without equals sign
			}
			key := strings.TrimSpace(line[:idx])
			value := ""
			if idx+1 < len(line) {
				value = strings.TrimSpace(line[idx+1:])
			}
			if len(value) > 1 {
				if (value[0] == '"' && value[len(value)-1] == '"') ||
					(value[0] == '\'' && value[len(value)-1] == '\'') {
					value = value[1 : len(value)-1]
				}
			}

			if err := os.Setenv(key, value); err != nil {
				return err
			}
		}
	}
	ClearCache()

	return nil
}

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
