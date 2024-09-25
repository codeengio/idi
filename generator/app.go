package generator

import (
	"embed"

	"github.com/codeengio/idi/filewriter"
	"github.com/rs/zerolog"
)

type App struct {
	Logger zerolog.Logger
}

func (f *App) GenerateNew(name, goModule string, templatesMap map[string]string, templateFS embed.FS) error {
	for fileName, templatePath := range templatesMap {
		f.Logger.Info().Str("create", fileName).Msg("generating file")
		err := filewriter.WriteFile(
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
