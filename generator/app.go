package generator

import (
	"io/fs"

	"github.com/codeengio/idi/writer"
	"github.com/rs/zerolog"
)

type App struct {
	logger zerolog.Logger
	writer writer.Writer
}

func NewApp(logger zerolog.Logger, w writer.Writer) *App {
	return &App{logger: logger, writer: w}
}

func (f *App) GenerateNew(name, goModule string, templatesMap map[string]string, templateFS fs.FS) error {
	for fileName, templatePath := range templatesMap {
		f.logger.Info().Str("create", fileName).Msg("generating file")
		err := f.writer.WriteTemplate(
			name,
			templatePath,
			fileName,
			templateFS,
			map[string]string{"AppName": name, "ModuleName": goModule},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
