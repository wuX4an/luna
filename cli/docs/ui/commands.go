package ui

import (
	"strings"
)

var commandMap = map[string]string{
	// INDEX
	"":             Manual.Index,
	"/index":       Manual.Index,
	"/usage":       Manual.Usage.Index,
	"/usage/start": Manual.Usage.Start,
	"u/start":      Manual.Usage.Start,

	// REFERENCE
	"/reference":        Manual.Reference.Index,
	"/r":                Manual.Reference.Index,
	"/reference/std":    Manual.Reference.Std.Index,
	"/std":              Manual.Reference.Std.Index,
	"/reference/cli":    Manual.Reference.Cli,
	"/cli":              Manual.Reference.Cli,
	"/reference/config": Manual.Reference.Config,
	"/config":           Manual.Reference.Config,

	// REFERENCE/STD
	"/reference/std/base64": Manual.Reference.Std.Base64,
	"/base64":               Manual.Reference.Std.Base64,
	"/std/base64":           Manual.Reference.Std.Base64,
	"/reference/std/random": Manual.Reference.Std.Random,
	"/random":               Manual.Reference.Std.Random,
	"/std/random":           Manual.Reference.Std.Random,
	"/reference/std/time":   Manual.Reference.Std.Time,
	"/time":                 Manual.Reference.Std.Time,
	"/std/time":             Manual.Reference.Std.Time,
	"/reference/std/tablex": Manual.Reference.Std.Tablex,
	"/tablex":               Manual.Reference.Std.Tablex,
	"/std/tablex":           Manual.Reference.Std.Tablex,
	"/reference/std/sqlite": Manual.Reference.Std.Sqlite,
	"/sqlite":               Manual.Reference.Std.Sqlite,
	"/std/sqlite":           Manual.Reference.Std.Sqlite,
	"/reference/std/test":   Manual.Reference.Std.Test,
	"/test":                 Manual.Reference.Std.Test,
	"/std/test":             Manual.Reference.Std.Test,
	"/reference/std/http":   Manual.Reference.Std.Http,
	"/http":                 Manual.Reference.Std.Http,
	"/std/http":             Manual.Reference.Std.Http,
	"/reference/std/math":   Manual.Reference.Std.Math,
	"/math":                 Manual.Reference.Std.Math,
	"/std/math":             Manual.Reference.Std.Math,
	"/reference/std/env":    Manual.Reference.Std.Env,
	"/env":                  Manual.Reference.Std.Env,
	"/std/env":              Manual.Reference.Std.Env,
	"/reference/std/crypto": Manual.Reference.Std.Crypto,
	"/crypto":               Manual.Reference.Std.Crypto,
	"/std/crypto":           Manual.Reference.Std.Crypto,
	"/reference/std/json":   Manual.Reference.Std.Json,
	"/json":                 Manual.Reference.Std.Json,
	"/std/json":             Manual.Reference.Std.Json,

	// WEB
	"/web":                     Manual.Reference.Web.Index,
	"/reference/web/animate":   Manual.Reference.Web.Animate,
	"/animate":                 Manual.Reference.Web.Animate,
	"/reference/web/audio":     Manual.Reference.Web.Audio,
	"/audio":                   Manual.Reference.Web.Audio,
	"/reference/web/await":     Manual.Reference.Web.Await,
	"/await":                   Manual.Reference.Web.Await,
	"/reference/web/clipboard": Manual.Reference.Web.Clipboard,
	"/clipboard":               Manual.Reference.Web.Clipboard,
	"/reference/web/console":   Manual.Reference.Web.Console,
	"/console":                 Manual.Reference.Web.Console,
	"/reference/web/cookie":    Manual.Reference.Web.Cookie,
	"/cookie":                  Manual.Reference.Web.Cookie,
	"/reference/web/dom":       Manual.Reference.Web.DOM,
	"/dom":                     Manual.Reference.Web.DOM,
	"/reference/web/fetch":     Manual.Reference.Web.Fetch,
	"/fetch":                   Manual.Reference.Web.Fetch,
	"/reference/web/keyboard":  Manual.Reference.Web.Keyboard,
	"/keyboard":                Manual.Reference.Web.Keyboard,
	"/reference/web/notify":    Manual.Reference.Web.Notify,
	"/notify":                  Manual.Reference.Web.Notify,
	"/reference/web/storage":   Manual.Reference.Web.Storage,
	"/storage":                 Manual.Reference.Web.Storage,
	"/reference/web/timer":     Manual.Reference.Web.Timer,
	"/timer":                   Manual.Reference.Web.Timer,
}

func processCommand(m Model) Model {
	inputVal := strings.ToLower(strings.TrimSpace(m.TextInput.Value()))
	if content, ok := commandMap[inputVal]; ok {
		m.Content = content
	} else {
		m.Content = "\u200B│ **Command not found:** `" + inputVal + "`\n\n" + Manual.Index
	}
	return m
}
