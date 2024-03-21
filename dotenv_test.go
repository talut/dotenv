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
