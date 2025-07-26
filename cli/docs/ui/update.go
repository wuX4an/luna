package ui

import (
	"luna/cli/docs/manual" // Importa tu paquete docs
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles incoming messages and updates the model accordingly.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
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
				inputVal := strings.ToLower(strings.TrimSpace(m.TextInput.Value()))
				switch inputVal {
				// case "/cli":
				// 	m.Content = manual.CliContent
				// case "/std":
				// 	m.Content = manual.StdContent
				case "/index", "":
					m.Content = manual.IndexContent
				default:
					m.Content = "\033[0;31mCommand not found: \033[0m'" +
						inputVal +
						"'\n" + manual.IndexContent
				}
				m.TextInput.SetValue("")
				m.Viewport.SetContent(m.Content)
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
				m.TextInput.SetValue("") // opcional: limpia input al entrar
			}
		// navegación solo si el input no está activo
		case "j", "down":
			if !m.InputActive {
				m.Viewport.ScrollDown(1)
			}
		case "k", "up":
			if !m.InputActive {
				m.Viewport.ScrollUp(1)
			}
		}
	case tea.WindowSizeMsg:
		if !m.Ready {
			m.Viewport = viewport.New(msg.Width, msg.Height-2)
			m.Viewport.SetContent(m.Content)
			m.Ready = true
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height - 2
		}
		m.TextInput.Width = msg.Width - 4
	}

	if m.InputActive {
		m.TextInput, cmd = m.TextInput.Update(msg)
	} else {
		m.Viewport, cmd = m.Viewport.Update(msg)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
