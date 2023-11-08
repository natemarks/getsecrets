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

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()
	workingDirectory, _ := os.Getwd()
	logger.Info().Msg("starting in " + workingDirectory)
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
		logger.Debug().Msgf("trying to get secret %s", value.SecretId)
		secretValue, err := secretsmanager.GetSecretValue(value.SecretId)
		if err != nil {
			logger.Error().Msgf("failed to get secret %s", value.SecretId)
			continue
		}
		logger.Debug().Msgf("successfully got secret %s", value.SecretId)
		absolutePath := workingDirectory + "/" + value.Filename
		logger.Debug().Msgf("trying to write secret %s to %s", value.SecretId, absolutePath)
		err = file.WriteFile(value.Filename, secretValue)
		if err != nil {
			logger.Error().Msgf("failed to write secret %s to %s", value.SecretId, absolutePath)
			logger.Error().Err(err)
			continue
		}
		logger.Debug().Msgf("successfully wrote secret %s to %s", value.SecretId, absolutePath)
	}
}
