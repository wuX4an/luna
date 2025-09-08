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
		m.Content = Manual.Usage.Index
	case "/usage/start", "u/start":
		m.Content = Manual.Usage.Start
	// REFERENCE
	case "/reference", "/r":
		m.Content = Manual.Reference.Index
	case "/reference/std", "/std":
		m.Content = Manual.Reference.Std.Index
	// REFERENCE/CLI
	case "/reference/cli", "/cli":
		m.Content = Manual.Reference.Cli
	// REFERENCE/CONFIG
	case "/reference/config", "/config":
		m.Content = Manual.Reference.Config
	// REFERENCE/STD
	case "/reference/std/base64", "/base64", "/std/base64":
		m.Content = Manual.Reference.Std.Base64
	case "/reference/std/random", "/random", "/std/random":
		m.Content = Manual.Reference.Std.Random
	case "/reference/std/time", "/time", "/std/time":
		m.Content = Manual.Reference.Std.Time
	case "/reference/std/tablex", "/tablex", "/std/tablex":
		m.Content = Manual.Reference.Std.Tablex
	case "/reference/std/sqlite", "/sqlite", "/std/sqlite":
		m.Content = Manual.Reference.Std.Sqlite
	case "/reference/std/test", "/test", "/std/test":
		m.Content = Manual.Reference.Std.Test
	case "/reference/std/http", "/http", "/std/http":
		m.Content = Manual.Reference.Std.Http
	case "/reference/std/math", "/math", "/std/math":
		m.Content = Manual.Reference.Std.Math
	case "/reference/std/env", "/env", "/std/env":
		m.Content = Manual.Reference.Std.Env
	case "/reference/std/crypto", "/crypto", "/std/crypto":
		m.Content = Manual.Reference.Std.Crypto
	case "/reference/std/json", "/json", "/std/json":
		m.Content = Manual.Reference.Std.Json
	default:
		content := Manual.Index
		m.Content = "\u200B│ **Command not found:** `" + inputVal + "`\n\n" + content
	}

	return m
}
