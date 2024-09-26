package main

import (
	"embed"
	"github.com/codeengio/idi/writer"
	"os"
	"time"

	"github.com/codeengio/idi/cmd"
	"github.com/codeengio/idi/generator"
	"github.com/rs/zerolog"
)

//go:embed all:templates
var templateFS embed.FS

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	appGen := generator.NewApp(log, writer.NewFS(log))
	appRunner := cmd.NewAppRunner(appGen, log, templateFS)
	rootCmd := appRunner.NewAppCmd()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to execute the command")
	}
}
