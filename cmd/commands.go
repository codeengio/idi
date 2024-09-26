package cmd

import (
	"errors"
	"fmt"
	"github.com/codeengio/idi/generator"
	"github.com/rs/zerolog"
	"io/fs"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var allowedArgs = []string{"new"}

type AppRunner struct {
	logger     zerolog.Logger
	templateFS fs.FS
	appGen     *generator.App
}

func NewAppRunner(appGen *generator.App, logger zerolog.Logger, templateFS fs.FS) *AppRunner {
	return &AppRunner{appGen: appGen, logger: logger, templateFS: templateFS}
}

func (a *AppRunner) NewAppCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "Creates a new app",
		Long:  `Creates a new app`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires 'new' arg")
			}

			if !slices.Contains(allowedArgs, args[0]) {
				return fmt.Errorf("the arg must be one of: %s", strings.Join(allowedArgs, ", "))
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := a.runNewApp(cmd, args)
			if err != nil {
				return
			}
		},
	}
	cmd.PersistentFlags().StringP("name", "n", "", "app name")
	cmd.PersistentFlags().StringP("module", "m", "", "Go module name (example.com/module/app)")

	return cmd
}

func (a *AppRunner) runNewApp(cmd *cobra.Command, _ []string) error {
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
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	module, err := cmd.Flags().GetString("module")
	if err != nil {
		return err
	}

	err = a.appGen.GenerateNew(name, module, templates, a.templateFS)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to generate new app")
		return err
	}

	return nil
}
