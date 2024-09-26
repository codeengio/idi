package generator

import (
	"errors"
	"github.com/codeengio/idi/mocks/github.com/codeengio/idi/writer"
	"github.com/codeengio/idi/tests"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestApp_GenerateNew(t *testing.T) {
	w := &writer.MockWriter{}

	w.On("WriteTemplate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	appGen := NewApp(zerolog.Logger{}, w)
	err := appGen.GenerateNew(
		"app",
		"example.com/module/app",
		map[string]string{"README.md": "templates/readme.md.tmpl"},
		tests.GetFS(),
	)
	if err != nil {
		t.Error(err)
	}

	if len(w.Calls) != 1 {
		t.Fatalf("expected 1 call, got=%d", len(w.Calls))
	}
}

func TestApp_GenerateNew_EmptyTemplates(t *testing.T) {
	w := &writer.MockWriter{}
	appGen := NewApp(zerolog.Logger{}, w)
	err := appGen.GenerateNew("app", "example.com/module/app", map[string]string{}, tests.GetFS())
	if err != nil {
		t.Error(err)
	}

	if len(w.Calls) > 0 {
		t.Fatalf("should not have any calls")
	}
}

func TestApp_GenerateNew_WriteError(t *testing.T) {
	w := &writer.MockWriter{}

	w.On("WriteTemplate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(errors.New("write error"))

	appGen := NewApp(zerolog.Logger{}, w)
	err := appGen.GenerateNew(
		"app",
		"example.com/module/app",
		map[string]string{"README.md": "templates/readme.md.tmpl"},
		tests.GetFS(),
	)
	if err == nil {
		t.Fatalf("should have errored")
	}
}
