package repl

import (
	"bytes"
	"fmt"
	"luna/std"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type LuaRepl struct {
	L   *lua.LState
	buf *bytes.Buffer
}

func NewLuaRepl() *LuaRepl {
	L := lua.NewState()
	// Registrar la std library
	std.RegisterAll(L)
	buf := &bytes.Buffer{}

	// Sobrescribir print para capturar salida
	L.SetGlobal("print", L.NewFunction(func(L *lua.LState) int {
		n := L.GetTop()
		parts := make([]string, n)
		for i := 1; i <= n; i++ {
			val := L.Get(i)
			if strVal, ok := val.(lua.LString); ok {
				// If it's a string, color it green
				parts[i-1] = "\033[3;90m" + string(strVal) + "\033[0m" // Green color
			} else {
				// Otherwise, use its default string representation
				parts[i-1] = L.ToStringMeta(val).String()
			}
		}
		// Output of print statements will be in a light gray/white (or green for strings)
		buf.WriteString(strings.Join(parts, "\t") + "\n")
		return 0
	}))

	return &LuaRepl{
		L:   L,
		buf: buf,
	}
}

func (r *LuaRepl) Eval(line string) string {
	r.buf.Reset()
	// Detectar si la línea define una variable local (simple heurística)
	trimmed := strings.TrimSpace(line)
	if strings.HasPrefix(trimmed, "local ") {
		// Enhanced warning color: Bold Yellow
		r.buf.WriteString("\033[1;33mWarning\033[0m: Local variable will not persist to subsequent lines.\n")
	}

	// Intentamos evaluar como expresión
	err := r.L.DoString("return " + line)
	if err != nil {
		// Si falla, intentamos ejecutarlo como statement normal
		err = r.L.DoString(line)
		if err != nil {
			// Error messages in Bright Red
			return r.buf.String() + "\033[91mError:\033[0m " + err.Error() + "\n"
		}
		return r.buf.String()
	}

	// Comprobar si la pila tiene resultado
	top := r.L.GetTop()
	if top == 0 {
		return r.buf.String()
	}

	ret := r.L.Get(top)
	r.L.Pop(1)

	var result string
	// For expression results, we don't know if they are strings from code,
	// so we'll keep them bright cyan as before.
	// If you want string *literals* specifically to be green when returned as an expression,
	// you'd need more complex parsing or Lua value inspection here.
	switch v := ret.(type) {
	case lua.LString:
		result = string(v)
	case lua.LNumber:
		result = fmt.Sprintf("%v", float64(v))
	case lua.LBool:
		if bool(v) {
			result = "true"
		} else {
			result = "false"
		}
	default:
		result = ret.String()
	}

	// Output of evaluated expressions will be in a light blue/cyan for distinctiveness
	if r.buf.Len() > 0 {
		return r.buf.String() + "\033[96m" + result + "\033[0m\n"
	}
	return "\033[96m" + result + "\033[0m\n"
}

func (r *LuaRepl) Close() {
	r.L.Close()
}
