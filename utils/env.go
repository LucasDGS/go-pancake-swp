package utils

import "os"

func GetEnv(key string, fallback string) string {
	// get value of env (string) and verify if this env exists
	envValue, exist := os.LookupEnv(key)

	// If exist return env
	if exist {
		return envValue
	}

	// If not exist return fallback
	return fallback
}
