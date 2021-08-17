package helpers

import (
	"os"
	"strconv"
	"strings"
)

// GetEnv returns value of env variable with a given key, or a given ifMissing value if env variable is not defined.
func GetEnv(key, ifMissing string) string {
	v, exists := os.LookupEnv(key)
	if !exists || v == "" {
		v = ifMissing
	}

	return v
}

// GetEnvInt64 returns value of env variable with a given key, or a given ifMissing value if env variable is not defined.
func GetEnvInt64(key string, ifMissing int64) int64 {
	v := GetEnv(key, "")
	if v == "" {
		return ifMissing
	}

	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return ifMissing
	}

	return n
}

// GetEnvBool returns value of env variable with a given key, or a given ifMissing value if env variable is not defined.
func GetEnvBool(key string, ifMissing bool) bool {
	v := GetEnv(key, "")
	if v == "" {
		return ifMissing
	}

	return IsTrue(v)
}

// IsTrue returns true if a given string is one of yes, true, 1.
func IsTrue(v string) bool {
	var trueValues = map[string]bool{
		"yes":  true,
		"true": true,
		"on":   true,
		"1":    true,
	}

	v = strings.TrimSpace(strings.ToLower(v))
	_, ok := trueValues[v]

	return ok
}
