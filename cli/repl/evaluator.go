package repl

import (
	"bytes"
	"fmt"
	"luna/src/luavm"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type LuaRepl struct {
	L   *lua.LState
	buf *bytes.Buffer
}

func NewLuaRepl() *LuaRepl {
	L := luavm.NewLuaVM()

	buf := &bytes.Buffer{}
	return &LuaRepl{
		L:   L,
		buf: buf,
	}
}

func (r *LuaRepl) Eval(line string) string {
	r.buf.Reset()
	trimmed := strings.TrimSpace(line)
	if strings.HasPrefix(trimmed, "local ") {
		r.buf.WriteString("\033[1;33mWarning\033[0m: Local variable will not persist to subsequent lines.\n")
	}

	err := r.L.DoString("return " + line)
	if err != nil {
		err = r.L.DoString(line)
		if err != nil {
			return r.buf.String() + "\033[91mError:\033[0m " + err.Error() + "\n"
		}
		r.L.SetTop(0) // <--- Limpia pila después de ejecutar statement
		return r.buf.String()
	}
	r.L.SetTop(0) // <--- Limpia pila si no hay resultados

	top := r.L.GetTop()
	if top == 0 {
		r.L.SetTop(0) // <--- Limpia pila si no hay resultados
		return r.buf.String()
	}

	ret := r.L.Get(top)
	r.L.Pop(1)

	var result string
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

	r.L.SetTop(0) // <--- Limpia pila después de obtener resultado

	if r.buf.Len() > 0 {
		return r.buf.String() + "\033[96m" + result + "\033[0m\n"
	}
	return "\033[96m" + result + "\033[0m\n"
}

func (r *LuaRepl) Close() {
	r.L.Close()
}
