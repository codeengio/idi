package tests

import (
	"embed"
	"io/fs"
)

//go:embed data/*
var data embed.FS

var GetFS = func() fs.FS {
	return data
}
