package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.BottomLeft = "┴"
		return titleStyle.BorderStyle(b)
	}()
	headerErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

// View renders the TUI.
func (m Model) View() string {
	if !m.Ready {
		return "Loading..."
	}

	contentView := m.Viewport.View()
	header := m.headerView()
	footer := m.footerView()

	// Si el input está activo, lo ponemos entre contenido y footer
	// if m.InputActive {
	// 	return fmt.Sprintf("%s\n%s\n%s\n%s", header, contentView, m.TextInput.View(), footer)
	// }
	//
	// Si no, solo contenido normal + footer
	return fmt.Sprintf("%s\n%s\n%s", header, contentView, footer)
}

func (m Model) headerView() string {
	title := m.HeaderMsg

	// Si no hay mensaje, usamos la primera línea del contenido
	if title == "" {
		lines := strings.SplitN(m.Content, "\n", 2)
		if len(lines) > 0 {
			title = lines[0]
		}
		title = strings.ReplaceAll(title, "#", "")
		title = strings.TrimSpace(title)
	}

	titleStyled := titleStyle.Render(title)
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(titleStyled)))
	return lipgloss.JoinHorizontal(lipgloss.Center, titleStyled, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))

	if m.InputActive {
		inputAboveLine := lipgloss.JoinVertical(lipgloss.Left,
			m.TextInput.View(),
			line,
		)

		return lipgloss.JoinHorizontal(lipgloss.Bottom, inputAboveLine, info)
	}

	return lipgloss.JoinHorizontal(lipgloss.Bottom, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
