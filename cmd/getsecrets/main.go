package main

import (
	"os"

	"github.com/natemarks/getsecrets/version"
	"github.com/rs/zerolog"
)

// makemine processes each argument FromUrl and FromFile, using the first one that works
// if there are no arguments it prints usage and exits with an error code
// there is no logging unless the debug logging is set
func main() {

	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()
	logger.Info().Msg("starting")

}
