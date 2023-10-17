package environment

import (
	"os"
	"strings"
)

// GetSecretsEnvVars return env var strings that start with GETSECRETS_
func GetSecretsEnvVars() (result []string) {
	all := os.Environ()
	for _, v := range all {
		if strings.HasPrefix(v, "GETSECRETS_") {
			result = append(result, v)
		}
	}

	return result
}

// GetSecretsEnvVarsMap returns a map of all the environment variables
// that start with GETSECRETS_ with the GETSECRETS_ prefix removed
func GetSecretsEnvVarsMap(vars []string) (result map[string]string) {
	result = make(map[string]string)
	var key, value string
	for _, v := range vars {
		if strings.HasPrefix(v, "GETSECRETS_") {
			parts := strings.Split(v, "=")
			value = strings.TrimPrefix(v, parts[0]+"=")
			key = strings.TrimPrefix(parts[0], "GETSECRETS_")
			result[key] = value
		}
	}

	return result
}
