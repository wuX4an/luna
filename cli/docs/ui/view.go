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
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

// View renders the TUI.
func (m Model) View() string {
	if !m.Ready {
		return "Loading..."
	}

	contentView := m.Viewport.View()

	if m.InputActive {
		return fmt.Sprintf("%s\n%s\n%s", m.headerView(), contentView, m.TextInput.View())
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), contentView, m.footerView())
}

func (m Model) headerView() string {
	lines := strings.SplitN(m.Content, "\n", 2)
	title := ""
	if len(lines) > 0 {
		title = lines[0]
	}
	title = strings.ReplaceAll(title, "#", "")
	title = strings.Replace(title, " ", "", 1)
	titleStyled := titleStyle.Render(title)
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(titleStyled)))
	return lipgloss.JoinHorizontal(lipgloss.Center, titleStyled, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
