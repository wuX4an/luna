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
	Json   string
}

type WEB struct {
	Index     string
	Animate   string
	Audio     string
	Await     string
	Clipboard string
	Console   string
	Cookie    string
	DOM       string
	Fetch     string
	Keyboard  string
	Notify    string
	Storage   string
	Timer     string
}

type Reference struct {
	Index  string
	Cli    string
	Lua    string
	Config string
	Std    STD
	Web    WEB
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
		Architecture: luna.MustReadDocFile("concepts/manual/architecture.md"),
	},
	Usage: Usage{
		Index: luna.MustReadDocFile("usage/index.md"),
		Start: luna.MustReadDocFile("usage/manual/start.md"),
	},
	Reference: Reference{
		Index:  luna.MustReadDocFile("reference/index.md"),
		Cli:    luna.MustReadDocFile("reference/cli/index.md"),
		Lua:    luna.MustReadDocFile("reference/lua/index.md"),
		Config: luna.MustReadDocFile("reference/config/index.md"),
		Std: STD{
			Index:  luna.MustReadDocFile("reference/std/index.md"),
			Base64: luna.MustReadDocFile("reference/std/manual/base64.md"),
			Random: luna.MustReadDocFile("reference/std/manual/random.md"),
			Time:   luna.MustReadDocFile("reference/std/manual/time.md"),
			Tablex: luna.MustReadDocFile("reference/std/manual/tablex.md"),
			Sqlite: luna.MustReadDocFile("reference/std/manual/sqlite.md"),
			Test:   luna.MustReadDocFile("reference/std/manual/test.md"),
			Http:   luna.MustReadDocFile("reference/std/manual/http.md"),
			Math:   luna.MustReadDocFile("reference/std/manual/math.md"),
			Env:    luna.MustReadDocFile("reference/std/manual/env.md"),
			Crypto: luna.MustReadDocFile("reference/std/manual/crypto.md"),
			Json:   luna.MustReadDocFile("reference/std/manual/json.md"),
		},
		Web: WEB{
			Index:     luna.MustReadDocFile("reference/std/web/index.md"),
			Animate:   luna.MustReadDocFile("reference/std/web/manual/animate.md"),
			Audio:     luna.MustReadDocFile("reference/std/web/manual/audio.md"),
			Await:     luna.MustReadDocFile("reference/std/web/manual/await.md"),
			Clipboard: luna.MustReadDocFile("reference/std/web/manual/clipboard.md"),
			Console:   luna.MustReadDocFile("reference/std/web/manual/console.md"),
			Cookie:    luna.MustReadDocFile("reference/std/web/manual/cookie.md"),
			DOM:       luna.MustReadDocFile("reference/std/web/manual/dom.md"),
			Fetch:     luna.MustReadDocFile("reference/std/web/manual/fetch.md"),
			Keyboard:  luna.MustReadDocFile("reference/std/web/manual/keyboard.md"),
			Notify:    luna.MustReadDocFile("reference/std/web/manual/notify.md"),
			Storage:   luna.MustReadDocFile("reference/std/web/manual/storage.md"),
			Timer:     luna.MustReadDocFile("reference/std/web/manual/timer.md"),
		},
	},
}
