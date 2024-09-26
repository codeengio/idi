package writer

import (
	"github.com/codeengio/idi/tests"
	"github.com/rs/zerolog"
	"os"
	"testing"
)

func TestFS_WriteTemplate(t *testing.T) {
	fs := NewFS(zerolog.Logger{})
	err := fs.WriteTemplate(
		"../out/app",
		"data/main.go.tmpl",
		"main.go",
		tests.GetFS(),
		map[string]string{"AppName": "app"},
	)
	if err != nil {
		t.Error(err)
	}
	stat, err := os.Stat("../out/app/main.go")
	if err != nil {
		t.Error(err)
	}
	if stat.Name() != "main.go" {
		t.Fatalf("expected file name=%s. got=%s", "main.go", stat.Name())
	}
}

func TestFS_WriteTemplate_NotExists(t *testing.T) {
	fs := NewFS(zerolog.Logger{})
	err := fs.WriteTemplate(
		"../out/app",
		"data/xyz.go.tmpl",
		"main.go",
		tests.GetFS(),
		map[string]string{"AppName": "app"},
	)
	if err == nil {
		t.Fatalf("expected error. got=nil")
	}
}
