package writer

import (
	"fmt"
	"github.com/rs/zerolog"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
)

type FS struct {
	logger zerolog.Logger
}

func NewFS(logger zerolog.Logger) *FS {
	return &FS{logger: logger}
}

func (fs *FS) WriteTemplate(appDir, templateName, outFileName string, templateFS fs.FS, args map[string]string) error {
	tmpl, err := template.ParseFS(templateFS, templateName)
	if err != nil {
		fs.logger.Error().Err(err).Msg("error parsing template")
		return err
	}

	fp := fmt.Sprintf("%s/%s", appDir, outFileName)
	if err := os.MkdirAll(filepath.Dir(fp), 0770); err != nil {
		fs.logger.Error().Err(err).Msg("error creating directory")
		return err
	}

	f, err := os.Create(fp)
	if err != nil {
		fs.logger.Error().Err(err).Msg("error creating file")
		return err
	}

	err = tmpl.Execute(f, args)
	if err != nil {
		fs.logger.Error().Err(err).Msg("error executing template")
		return err
	}

	return err
}
