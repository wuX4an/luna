package ui

// View renders the TUI.
func (m Model) View() string {
	if !m.Ready {
		return "Loading..."
	}

	// Display viewport content above and text input below
	return m.Viewport.View() + "\n" + m.TextInput.View()
}
