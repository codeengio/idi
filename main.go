package main

import (
	"embed"
	"os"
	"time"

	"github.com/codeengio/idi/cmd"
	"github.com/codeengio/idi/generator"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

//go:embed all:templates
var templateFS embed.FS

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	rootCmd := cmd.NewRootCmd(runNewApp(log))
	rootCmd.PersistentFlags().StringP("name", "n", "", "app name")
	rootCmd.PersistentFlags().StringP("module", "m", "", "Go module name (example.com/module/app)")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to execute the command")
	}
}

func runNewApp(logger zerolog.Logger) func(*cobra.Command, []string) error {
	templates := map[string]string{
		"README.md":              "templates/readme.md.tmpl",
		"main.go":                "templates/main.go.tmpl",
		"go.mod":                 "templates/go.mod.tmpl",
		"doc.go":                 "templates/doc.go.tmpl",
		"config/config.go":       "templates/config/config.go.tmpl",
		"db/db.go":               "templates/db/db.go.tmpl",
		"db/migration.go":        "templates/db/migration.go.tmpl",
		"db/migrations/init.sql": "templates/db/migrations/init.sql.tmpl",
		"internal/api/router.go": "templates/internal/api/router.go.tmpl",
		"docker-compose.yml":     "templates/docker-compose.yml.tmpl",
		"Makefile":               "templates/makefile.tmpl",
		".gitignore":             "templates/gitignore.tmpl",
		".env":                   "templates/env.tmpl",
		"pkg/hanko/auth.go":      "templates/pkg/hanko/auth.go.tmpl",
	}

	return func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		module, err := cmd.Flags().GetString("module")
		if err != nil {
			return err
		}

		appGen := generator.App{Logger: logger}
		err = appGen.GenerateNew(name, module, templates, templateFS)
		if err != nil {
			logger.Error().Err(err).Msg("failed to generate new app")
			return err
		}

		return nil
	}
}
