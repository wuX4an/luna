package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Tui initializes and starts the Bubble Tea program.
// (This is an alternative if you prefer to encapsulate the program start.)

func Tui() {
	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
