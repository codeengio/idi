package filewriter

import (
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
)

func WriteFile(appName, templateName, fileName string, args map[string]string) error {
	tmpl, err := template.ParseFiles(templateName)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	fp := fmt.Sprintf("%s/%s", appName, fileName)
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
