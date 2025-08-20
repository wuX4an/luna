package ui

import (
	"luna"
)

// Subestructuras
type Concepts struct {
	Index        string
	Architecture string
}

type STD struct {
	Index  string
	Base64 string
	Random string
	Time   string
	Tablex string
	Sqlite string
	Test   string
	Http   string
	Math   string
	Env    string
	Crypto string
}

type Reference struct {
	Index  string
	Cli    string
	Lua    string
	Config string
	Std    STD
}
type Usage struct {
	Index string
	Start string
}

// ManualStruct contiene todo el contenido embebido
type ManualStruct struct {
	Index     string
	Concepts  Concepts
	Usage     Usage
	Reference Reference
}

// Manual es la instancia global lista para usar
var Manual = ManualStruct{
	Index: luna.MustReadDocFile("index.md"),
	Concepts: Concepts{
		Index:        luna.MustReadDocFile("concepts/index.md"),
		Architecture: luna.MustReadDocFile("concepts/architecture.md"),
	},
	Usage: Usage{
		Index: luna.MustReadDocFile("usage/index.md"),
		Start: luna.MustReadDocFile("usage/start.md"),
	},
	Reference: Reference{
		Index:  luna.MustReadDocFile("reference/index.md"),
		Cli:    luna.MustReadDocFile("reference/cli/index.md"),
		Lua:    luna.MustReadDocFile("reference/lua/index.md"),
		Config: luna.MustReadDocFile("reference/config/index.md"),
		Std: STD{
			Index:  luna.MustReadDocFile("reference/std/index.md"),
			Base64: luna.MustReadDocFile("reference/std/base64.md"),
			Random: luna.MustReadDocFile("reference/std/random.md"),
			Time:   luna.MustReadDocFile("reference/std/time.md"),
			Tablex: luna.MustReadDocFile("reference/std/tablex.md"),
			Sqlite: luna.MustReadDocFile("reference/std/sqlite.md"),
			Test:   luna.MustReadDocFile("reference/std/test.md"),
			Http:   luna.MustReadDocFile("reference/std/http.md"),
			Math:   luna.MustReadDocFile("reference/std/math.md"),
			Env:    luna.MustReadDocFile("reference/std/env.md"),
			Crypto: luna.MustReadDocFile("reference/std/crypto.md"),
		},
	},
}
