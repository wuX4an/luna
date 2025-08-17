package ui

import (
	"luna"
)

// Subestructuras
type Concepts struct {
	Index        string
	Architecture string
	Workflow     string
}

type CLI struct {
	Index    string
	Commands string
	Flags    string
}

type STD struct {
	Index  string
	Base64 string
	Random string
	Time   string
	Tablex string
	Sqlite string
}

type Reference struct {
	Index string
	CLI   CLI
	Lua   string
	Std   STD
}

// ManualStruct contiene todo el contenido embebido
type ManualStruct struct {
	Index     string
	Concepts  Concepts
	Usage     string
	Reference Reference
}

// Manual es la instancia global lista para usar
var Manual = ManualStruct{
	Index: luna.MustReadDocFile("index.md"),
	Concepts: Concepts{
		Index:        luna.MustReadDocFile("concepts/index.md"),
		Architecture: luna.MustReadDocFile("concepts/architecture.md"),
		Workflow:     luna.MustReadDocFile("concepts/workflow.md"),
	},
	Usage: luna.MustReadDocFile("usage/index.md"),
	Reference: Reference{
		Index: luna.MustReadDocFile("reference/index.md"),
		CLI: CLI{
			Index:    luna.MustReadDocFile("reference/cli/index.md"),
			Commands: luna.MustReadDocFile("reference/cli/commands.md"),
			Flags:    luna.MustReadDocFile("reference/cli/flags.md"),
		},
		Lua: luna.MustReadDocFile("reference/lua/index.md"),
		Std: STD{
			Index:  luna.MustReadDocFile("reference/std/index.md"),
			Base64: luna.MustReadDocFile("reference/std/base64.md"),
			Random: luna.MustReadDocFile("reference/std/random.md"),
			Time:   luna.MustReadDocFile("reference/std/time.md"),
			Tablex: luna.MustReadDocFile("reference/std/tablex.md"),
			Sqlite: luna.MustReadDocFile("reference/std/sqlite.md"),
		},
	},
}
