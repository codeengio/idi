package filewriter

import (
	"embed"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
)

func WriteFile(appName, templateName, outFileName string, templateFS embed.FS, args map[string]string) error {
	tmpl, err := template.ParseFS(templateFS, templateName)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	fp := fmt.Sprintf("%s/%s", appName, outFileName)
	if err := os.MkdirAll(filepath.Dir(fp), 0770); err != nil {
		return err
	}

	f, err := os.Create(fp)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = tmpl.Execute(f, args)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return err
}
