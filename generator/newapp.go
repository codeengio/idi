package generator

import (
	"embed"
	"fmt"
	"log/slog"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/codeengio/idi/filewriter"
)

type (
	errMsg error
)

const (
	newAppTitle = "What's the name of your app?"
	moduleTitle = "What's the name of your Go module?"
)

type model struct {
	title        string
	textInput    textinput.Model
	appName      string
	moduleName   string
	templatesMap map[string]string
	templateFS   embed.FS
	err          error
}

func NewAppInitialModel(templatesMap map[string]string, templateFS embed.FS) model {
	ti := textinput.New()
	ti.Placeholder = "my-app"
	ti.Focus()
	ti.CharLimit = 1000
	ti.Width = 100

	return model{
		title:        newAppTitle,
		textInput:    ti,
		templatesMap: templatesMap,
		templateFS:   templateFS,
		err:          nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.title == newAppTitle {
				m.appName = m.textInput.Value()
			}
			if m.title == moduleTitle {
				m.moduleName = m.textInput.Value()
			}

			m.title = moduleTitle
			m.textInput.Placeholder = fmt.Sprintf("example.com/my-module/%s", m.appName)
			m.textInput.Reset()

			if m.appName != "" && m.moduleName != "" {
				err := m.WriteAllFiles()
				if err != nil {
					slog.Error(err.Error())
				}
				return m, tea.Quit
			}

			return m, cmd
		}

	case errMsg:
		slog.Error(msg.Error())
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.title,
		m.textInput.View(),
		"(esc to quit | ctrl+c to quit)",
	) + "\n"
}

func (m model) WriteAllFiles() error {
	for fileName, templatePath := range m.templatesMap {
		err := filewriter.WriteFile(
			m.appName,
			templatePath,
			fileName,
			m.templateFS,
			map[string]string{"AppName": m.appName, "ModuleName": m.moduleName},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
