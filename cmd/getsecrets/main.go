package main

import (
	"os"

	"github.com/natemarks/getsecrets/secretsmanager"

	"github.com/natemarks/getsecrets/config"
	"github.com/natemarks/getsecrets/environment"
	"github.com/natemarks/getsecrets/file"
	"github.com/natemarks/getsecrets/version"
	"github.com/rs/zerolog"
)

func WriteSecretToFile(secret config.Secret) (err error) {
	secretValue, err := secretsmanager.GetSecretValue(secret.SecretId)
	if err != nil {
		return err
	}
	err = file.WriteFile(secret.Filename, secretValue)
	return err
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()
	logger.Info().Msg("starting")
	// get all the environment variables that start with GETSECRETS_
	envvars := environment.GetSecretsEnvVars()
	for _, v := range envvars {
		logger.Debug().Msgf("found env var %s", v)
	}
	// organize then env vars into a map pf related keys
	envvarsMap := environment.GetSecretsEnvVarsMap(envvars)
	// get the config and secrets from the env vars
	config, secrets := config.GetConfig(envvarsMap)
	// if debug is set, set the global log level to debug
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// debug logging of the input
	for key, value := range secrets {
		logger.Debug().Msgf(
			"will lookup %s for %s and write it to %s",
			value.SecretId,
			key,
			value.Filename)
	}
	for _, value := range secrets {
		err := WriteSecretToFile(value)
		if err != nil {
			logger.Error().Err(err).Msgf(
				"error getting %s - %s to %s : %v",
				value.Name,
				value.SecretId,
				value.Filename,
				err)
			continue
		}
		logger.Info().Msgf(
			"wrote %s - %s to %s",
			value.Name,
			value.SecretId,
			value.Filename)
	}
}
