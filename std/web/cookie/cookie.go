//go:build js && wasm

package cookie

import (
	"net/url"
	"strconv"
	"strings"
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// ---------------------------
// Loader del módulo
// ---------------------------
func Loader(L *lua.LState) *lua.LTable {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"set":    Set,
		"get":    Get,
		"delete": Delete,
		"list":   List,
	})
	return mod
}

// ---------------------------
// Funciones
// ---------------------------

// cookie.set(name, value, options)
func Set(L *lua.LState) int {
	name := L.CheckString(1)
	value := L.CheckString(2)

	// Opciones
	options := L.OptTable(3, L.NewTable())
	var cookieStr strings.Builder
	cookieStr.WriteString(url.QueryEscape(name))
	cookieStr.WriteString("=")
	cookieStr.WriteString(url.QueryEscape(value))

	if expires := options.RawGetString("expires"); expires != lua.LNil {
		if secs, ok := expires.(lua.LNumber); ok {
			cookieStr.WriteString("; max-age=")
			cookieStr.WriteString(strconv.Itoa(int(secs)))
		}
	}
	if path := options.RawGetString("path"); path != lua.LNil {
		cookieStr.WriteString("; path=")
		cookieStr.WriteString(path.String())
	}
	if domain := options.RawGetString("domain"); domain != lua.LNil {
		cookieStr.WriteString("; domain=")
		cookieStr.WriteString(domain.String())
	}
	if secure := options.RawGetString("secure"); secure == lua.LTrue {
		cookieStr.WriteString("; secure")
	}

	js.Global().Get("document").Set("cookie", cookieStr.String())
	return 0
}

// cookie.get(name)
func Get(L *lua.LState) int {
	name := url.QueryEscape(L.CheckString(1))
	cookies := strings.Split(js.Global().Get("document").Get("cookie").String(), ";")
	for _, c := range cookies {
		c = strings.TrimSpace(c)
		parts := strings.SplitN(c, "=", 2)
		if len(parts) == 2 && parts[0] == name {
			val, _ := url.QueryUnescape(parts[1])
			L.Push(lua.LString(val))
			return 1
		}
	}
	L.Push(lua.LNil)
	return 1
}

// cookie.delete(name)
func Delete(L *lua.LState) int {
	name := L.CheckString(1)
	// Path opcional
	path := L.OptString(2, "/")

	// Construir cookie para expirar inmediatamente
	cookieStr := url.QueryEscape(name) + "=; max-age=0; path=" + path

	// Opcional: también se puede forzar expires en el pasado
	cookieStr += "; expires=Thu, 01 Jan 1970 00:00:00 GMT"

	// Escribir directamente en document.cookie
	js.Global().Get("document").Set("cookie", cookieStr)

	return 0
}

// cookie.list()
func List(L *lua.LState) int {
	tbl := L.NewTable()
	cookies := strings.Split(js.Global().Get("document").Get("cookie").String(), ";")
	for _, c := range cookies {
		c = strings.TrimSpace(c)
		parts := strings.SplitN(c, "=", 2)
		if len(parts) == 2 {
			key, _ := url.QueryUnescape(parts[0])
			val, _ := url.QueryUnescape(parts[1])
			tbl.RawSetString(key, lua.LString(val))
		}
	}
	L.Push(tbl)
	return 1
}
