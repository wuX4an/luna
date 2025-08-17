package ui

import (
	"strings"
)

func processCommand(m Model) Model {
	inputVal := strings.ToLower(strings.TrimSpace(m.TextInput.Value()))

	switch inputVal {
	// INDEX
	case "", "/index":
		m.Content = Manual.Index
	// USAGE
	case "/usage":
		m.Content = Manual.Usage
	// REFERENCE
	case "/reference", "/r":
		m.Content = Manual.Reference.Index
	case "/reference/std", "/std":
		m.Content = Manual.Reference.Std.Index
	case "/reference/std/base64", "/base64", "/std/base64":
		m.Content = Manual.Reference.Std.Base64
	default:
		content := Manual.Index
		m.Content = "\u200B│ **Command not found:** `" + inputVal + "`\n\n" + content
	}

	return m
}
