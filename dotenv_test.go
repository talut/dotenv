package dotenv

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func setupEnv(key, value string) func() {
	os.Setenv(key, value)
	return func() {
		os.Unsetenv(key)
	}
}

func TestLoad(t *testing.T) {
	// Create temporary .env files for testing
	defaultEnvContent := []byte("DEFAULT_VAR=default_value\nSHARED_VAR=from_default")
	customEnvContent := []byte("CUSTOM_VAR=custom_value\nSHARED_VAR=from_custom")

	// Write default .env file
	err := os.WriteFile(".env", defaultEnvContent, 0644)
	require.NoError(t, err)
	defer os.Remove(".env")

	// Write custom .env file
	customEnvPath := ".env.custom"
	err = os.WriteFile(customEnvPath, customEnvContent, 0644)
	require.NoError(t, err)
	defer os.Remove(customEnvPath)

	// Test 1: Load default .env
	t.Run("Load default .env file", func(t *testing.T) {
		// Clear environment and cache first
		os.Unsetenv("DEFAULT_VAR")
		os.Unsetenv("SHARED_VAR")
		ClearCache()

		err := Load()
		require.NoError(t, err)

		// Check environment variables are set
		val := os.Getenv("DEFAULT_VAR")
		require.Equal(t, "default_value", val)

		sharedVal := os.Getenv("SHARED_VAR")
		require.Equal(t, "from_default", sharedVal)

		// Check cache is populated
		cachedVal, exists := cache["DEFAULT_VAR"]
		require.True(t, exists)
		require.Equal(t, "default_value", cachedVal)
	})

	// Test 2: Load custom .env file
	t.Run("Load custom .env file", func(t *testing.T) {
		// Clear environment and cache first
		os.Unsetenv("DEFAULT_VAR")
		os.Unsetenv("CUSTOM_VAR")
		os.Unsetenv("SHARED_VAR")
		ClearCache()

		err := Load(customEnvPath)
		require.NoError(t, err)

		// Check custom environment variable
		val := os.Getenv("CUSTOM_VAR")
		require.Equal(t, "custom_value", val)

		// Check default env var doesn't exist
		defaultVal := os.Getenv("DEFAULT_VAR")
		require.Equal(t, "", defaultVal)

		// Check cache is populated correctly
		cachedVal, exists := cache["CUSTOM_VAR"]
		require.True(t, exists)
		require.Equal(t, "custom_value", cachedVal)
	})

	// Test 3: Load multiple .env files
	t.Run("Load multiple .env files", func(t *testing.T) {
		// Clear environment and cache first
		os.Unsetenv("DEFAULT_VAR")
		os.Unsetenv("CUSTOM_VAR")
		os.Unsetenv("SHARED_VAR")
		ClearCache()

		err := Load(".env", customEnvPath)
		require.NoError(t, err)

		// Both files' unique vars should be loaded
		defaultVal := os.Getenv("DEFAULT_VAR")
		require.Equal(t, "default_value", defaultVal)

		customVal := os.Getenv("CUSTOM_VAR")
		require.Equal(t, "custom_value", customVal)

		// Shared var should have value from last file loaded
		sharedVal := os.Getenv("SHARED_VAR")
		require.Equal(t, "from_custom", sharedVal)

		// Check cache is populated correctly
		cachedDefault, exists := cache["DEFAULT_VAR"]
		require.True(t, exists)
		require.Equal(t, "default_value", cachedDefault)

		cachedShared, exists := cache["SHARED_VAR"]
		require.True(t, exists)
		require.Equal(t, "from_custom", cachedShared)
	})

	// Test 4: Non-existent file
	t.Run("Non-existent file", func(t *testing.T) {
		ClearCache()
		err := Load("non_existent_file.env")
		require.NoError(t, err) // Should not error for non-existent files
	})

	// Test 5: Test quoting in .env file
	t.Run("Test quotes in .env file", func(t *testing.T) {
		quotedEnvContent := []byte(`
QUOTED_DOUBLE="double quoted value"
QUOTED_SINGLE='single quoted value'
QUOTED_NESTED="'nested quotes'"
`)
		quotedEnvPath := ".env.quoted"
		err = os.WriteFile(quotedEnvPath, quotedEnvContent, 0644)
		require.NoError(t, err)
		defer os.Remove(quotedEnvPath)

		ClearCache()
		err := Load(quotedEnvPath)
		require.NoError(t, err)

		// Check that quotes are properly removed
		doubleVal := os.Getenv("QUOTED_DOUBLE")
		require.Equal(t, "double quoted value", doubleVal)

		singleVal := os.Getenv("QUOTED_SINGLE")
		require.Equal(t, "single quoted value", singleVal)

		nestedVal := os.Getenv("QUOTED_NESTED")
		require.Equal(t, "'nested quotes'", nestedVal)

		// Check cache
		cachedVal, exists := cache["QUOTED_DOUBLE"]
		require.True(t, exists)
		require.Equal(t, "double quoted value", cachedVal)
	})
}

func TestFunctions(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{
			name: "GetInt and MustGetInt",
			testFunc: func(t *testing.T) {
				defer setupEnv("TEST_INT", "123")()

				val := GetInt("TEST_INT", 0)
				require.Equal(t, 123, val)

				val = GetInt("NON_EXISTING_KEY", 0)
				require.Equal(t, 0, val)

				defer setupEnv("INVALID_INT", "abc")()
				val = GetInt("INVALID_INT", 0)
				require.Equal(t, 0, val)

				val2 := MustGetInt("TEST_INT")
				require.Equal(t, 123, val2)

				require.Panics(t, func() { MustGetInt("NON_EXISTING_KEY") })
				require.Panics(t, func() { MustGetInt("INVALID_INT") })
			},
		},
		{
			name: "GetBool and MustGetBool",
			testFunc: func(t *testing.T) {
				defer setupEnv("TEST_BOOL", "true")()

				val := GetBool("TEST_BOOL", false)
				require.Equal(t, true, val)

				val = GetBool("NON_EXISTING_KEY", false)
				require.Equal(t, false, val)

				defer setupEnv("INVALID_BOOL", "abc")()
				val = GetBool("INVALID_BOOL", false)
				require.Equal(t, false, val)

				val2 := MustGetBool("TEST_BOOL")
				require.Equal(t, true, val2)

				require.Panics(t, func() { MustGetBool("NON_EXISTING_KEY") })
				require.Panics(t, func() { MustGetBool("INVALID_BOOL") })
			},
		},
		{
			name: "GetFloat and MustGetFloat",
			testFunc: func(t *testing.T) {
				defer setupEnv("TEST_FLOAT", "123.45")()

				val := GetFloat("TEST_FLOAT", 0)
				require.Equal(t, 123.45, val)

				val = GetFloat("NON_EXISTING_KEY", 0)
				require.Equal(t, 0.0, val)

				defer setupEnv("INVALID_FLOAT", "abc")()
				val = GetFloat("INVALID_FLOAT", 0)
				require.Equal(t, 0.0, val)

				val2 := MustGetFloat("TEST_FLOAT")
				require.Equal(t, 123.45, val2)

				require.Panics(t, func() { MustGetFloat("NON_EXISTING_KEY") })
				require.Panics(t, func() { MustGetFloat("INVALID_FLOAT") })
			},
		},
		{
			name: "GetDuration and MustGetDuration",
			testFunc: func(t *testing.T) {
				defer setupEnv("TEST_DURATION", "1s")()

				val := GetDuration("TEST_DURATION", 0)
				require.Equal(t, "1s", val.String())

				val = GetDuration("NON_EXISTING_KEY", 0)
				require.Equal(t, "0s", val.String())

				defer setupEnv("INVALID_DURATION", "abc")()
				val = GetDuration("INVALID_DURATION", 0)
				require.Equal(t, "0s", val.String())

				val2 := MustGetDuration("TEST_DURATION")
				require.Equal(t, "1s", val2.String())

				require.Panics(t, func() { MustGetDuration("NON_EXISTING_KEY") })
				require.Panics(t, func() { MustGetDuration("INVALID_DURATION") })
			},
		},
		{
			name: "GetString and MustGetString",
			testFunc: func(t *testing.T) {
				defer setupEnv("TEST_STRING", "test")()

				val := GetString("TEST_STRING", "")
				require.Equal(t, "test", val)

				val = GetString("NON_EXISTING_KEY", "")
				require.Equal(t, "", val)

				val2 := MustGetString("TEST_STRING")
				require.Equal(t, "test", val2)

				require.Panics(t, func() { MustGetString("NON_EXISTING_KEY") })
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}

func TestClearCache(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "test_value")
	val := GetString("TEST_ENV_VAR", "")
	require.Equal(t, "test_value", val)
	ClearCache()
	os.Setenv("TEST_ENV_VAR", "new_value")
	val = GetString("TEST_ENV_VAR", "")
	require.Equal(t, "new_value", val)
	os.Unsetenv("TEST_ENV_VAR")
}
