package ui

import (
	// Importa tu paquete docs
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

// Update handles incoming messages and updates the model accordingly.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	// Teclas
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+d":
			return m, tea.Quit

		case "q":
			if !m.InputActive {
				return m, tea.Quit
			}

		case "enter":
			if m.InputActive {
				m = processCommand(m)
				m.TextInput.SetValue("")
				m.Viewport.SetContent(renderMarkdown(m.Content))
				m.Viewport.GotoTop()
				m.InputActive = false
				m.TextInput.Blur()
			}

		case "esc":
			if m.InputActive {
				m.InputActive = false
				m.TextInput.Blur()
			}

		case "/":
			if !m.InputActive {
				m.InputActive = true
				m.TextInput.Focus()
				m.TextInput.SetValue("")
			}

		// Navegación
		case "j", "down":
			if !m.InputActive {
				m.Viewport.ScrollDown(1)
			}
		case "k", "up":
			if !m.InputActive {
				m.Viewport.ScrollUp(1)
			}
		}

	// Redimensionamiento
	case tea.WindowSizeMsg:
		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height-2)
			m.Viewport.SetContent(renderMarkdown(m.Content))
			m.Ready = true
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height - 2
		}
		m.TextInput.Width = msg.Width - 4
	}

	// Actualiza input o viewport según corresponda
	if m.InputActive {
		m.TextInput, cmd = m.TextInput.Update(msg)
	} else {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func renderMarkdown(input string) string {
	out, err := glamour.Render(input, "dark") // puedes usar "light", "dracula", etc
	if err != nil {
		return input // fallback a texto plano si falla
	}
	return strings.TrimPrefix(out, "\n")
}
