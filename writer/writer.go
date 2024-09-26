package writer

import "io/fs"

type Writer interface {
	WriteTemplate(appDir, templateName, outFileName string, templateFS fs.FS, args map[string]string) error
}
