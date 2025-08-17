package ui

import (
	"luna"
	"strings"
)

func processCommand(m Model) Model {
	inputVal := strings.ToLower(strings.TrimSpace(m.TextInput.Value()))

	switch inputVal {
	case "", "/index":
		m.Content, _ = luna.ReadDocFile("index.md")
		// case "/cli":
		//     m.Content = manual.CliContent
		// case "/std":
		//     m.Content = manual.StdContent
	default:
		content, err := luna.ReadDocFile("index.md")
		if err != nil {
			content = "Error loading index.md"
		}
		m.Content = "\u200B│ **Command not found:** `" + inputVal + "`\n\n" + content
	}

	return m
}
