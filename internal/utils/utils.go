package utils

import "os"

// AppGetEnv gets the value of an environment variable if it exists, otherwise returns the passed defaultValue
func AppGetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
