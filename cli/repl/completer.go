package repl

import (
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type LuaAutoCompleter struct {
	repl *LuaRepl
}

func (r *LuaRepl) Completer(line string) []string {
	var completions []string
	// Palabras clave comunes en Lua para autocompletar
	keywords := []string{}

	// Añadir keywords que empiecen con line
	for _, kw := range keywords {
		if strings.HasPrefix(kw, line) {
			completions = append(completions, kw)
		}
	}

	// Iterar globals en el estado Lua
	globals := r.L.G.Global
	globals.ForEach(func(key, val lua.LValue) {
		keyStr := key.String()
		if strings.HasPrefix(keyStr, line) {
			completions = append(completions, keyStr)
		}
	})

	return completions
}
