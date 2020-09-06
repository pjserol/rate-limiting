package config

import (
	"os"
	"strconv"
	"strings"
)

func getEnvInt(variable string, def int) int {
	tmp := os.Getenv(variable)

	if tmp != "" {
		val, err := strconv.Atoi(tmp)

		if err == nil {
			return val
		}
	}

	return def
}

func getEnvString(variable string, def string) string {
	val := os.Getenv(variable)

	if val != "" {
		return val
	}

	return def
}

func getEnvBool(variable string, def bool) bool {
	val := os.Getenv(variable)
	val = strings.ToLower(strings.TrimSpace(val))
	switch val {
	case "true":
		return true
	case "false":
		return false
	default:
		return def
	}
}
