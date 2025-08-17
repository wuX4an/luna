package ui

import (
	"strings"
)

func processCommand(m Model) Model {
	inputVal := strings.ToLower(strings.TrimSpace(m.TextInput.Value()))

	switch inputVal {
	case "", "/index":
		m.Content = Manual.Index
	// case "/cli":
	//     m.Content = manual.CliContent
	case "/usage":
		m.Content = Manual.Usage
	default:
		content := Manual.Index
		m.Content = "\u200B│ **Command not found:** `" + inputVal + "`\n\n" + content
	}

	return m
}
