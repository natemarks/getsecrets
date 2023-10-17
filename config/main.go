package config

import (
	"fmt"
	"strings"
)

type Secret struct {
	Name     string `json:"name"`
	SecretId string `json:"secretId"`
	Filename string `json:"filename"`
	Database string `json:"database"`
}

type Config struct {
	Debug bool `json:"debug"`
}

func updateSecret(secret Secret, key, value string) Secret {
	switch key {
	case "SECRETID":
		secret.SecretId = value
	case "FILENAME":
		secret.Filename = value
	case "DATABASE":
		secret.Database = value
	default:
		panic(fmt.Sprintf("unknown key %s", key))
	}
	return secret
}
func GetConfig(varMap map[string]string) (config Config, secrets map[string]Secret, err error) {
	secrets = make(map[string]Secret)
	config.Debug = false
	for key, value := range varMap {
		if key == "DEBUG" {
			if strings.ToLower(value) == "true" {
				config.Debug = true
				continue
			}
		}
		parts := strings.Split(key, "_")
		name := parts[0]

		val, ok := secrets[parts[0]]
		if !ok {
			thisSecret := Secret{
				Name: name,
			}
			secrets[name] = updateSecret(thisSecret, parts[1], value)
			continue
		}
		secrets[name] = updateSecret(val, parts[1], value)
	}
	return config, secrets, nil
}
