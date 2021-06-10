package env

import (
	"os"
	"strconv"
)

// String return env string
func String(key string) string {
	return os.Getenv(key)
}

// Bool return env bool
func Bool(key string) bool {
	val := os.Getenv(key)
	if val == "" {
		return false
	}

	ret, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}

	return ret
}

// Int return env int
func Int(env string) int {
	tm := os.Getenv(env)
	if tm == "" {
		return 0
	}

	i, err := strconv.Atoi(tm)
	if err != nil {
		return 0
	}

	return i
}
