package main

import (
	"embed"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/codeengio/idi/cmd"
	"github.com/codeengio/idi/generator"
	"github.com/spf13/cobra"
)

//go:embed all:templates
var templateFS embed.FS

func main() {
	rootCmd := cmd.NewRootCmd(runNewApp)
	rootCmd.Execute()
}

func runNewApp(cmd *cobra.Command, args []string) error {
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
	}
	p := tea.NewProgram(generator.NewAppInitialModel(templates, templateFS))
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
