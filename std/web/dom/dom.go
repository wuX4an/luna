//go:build js && wasm
// +build js,wasm

package dom

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

type DOMUserData struct {
	Elem js.Value
}

const domTypeName = "DOMElement"

// ---------------------------
// Métodos de DOMElement
// ---------------------------
func registerMethods(L *lua.LState) *lua.LTable {
	mt := L.NewTypeMetatable(domTypeName)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		// ---------------------------
		// Manipulación de hijos
		// ---------------------------
		"append": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			child := checkDOM(L, 2)
			ud.Elem.Call("appendChild", child.Elem)
			return 0
		},
		"appendText": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			text := L.CheckString(2)
			current := ud.Elem.Get("textContent").String()
			ud.Elem.Set("textContent", current+text)
			return 0
		},
		"clear": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			for child := ud.Elem.Get("firstChild"); child.Truthy(); child = ud.Elem.Get("firstChild") {
				ud.Elem.Call("removeChild", child)
			}
			return 0
		},
		"remove": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			parent := ud.Elem.Get("parentNode")
			if parent.Truthy() {
				parent.Call("removeChild", ud.Elem)
			}
			return 0
		},
		"replaceWith": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			newUD := checkDOM(L, 2)
			parent := ud.Elem.Get("parentNode")
			if parent.Truthy() {
				parent.Call("replaceChild", newUD.Elem, ud.Elem)
			}
			return 0
		},

		// ---------------------------
		// Visibilidad
		// ---------------------------
		"hide": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			ud.Elem.Get("style").Set("display", "none")
			return 0
		},
		"show": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			ud.Elem.Get("style").Set("display", "block") // o "flex" según layout
			return 0
		},
		"toggle": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			style := ud.Elem.Get("style")
			current := style.Get("display").String()

			if current == "none" {
				// Restaurar display original si se guardó
				orig := ud.Elem.Get("dataset").Get("originalDisplay").String()
				if orig == "" {
					orig = "block" // valor por defecto
				}
				style.Set("display", orig)
			} else {
				// Guardar display original antes de ocultar
				if style.Get("display").Truthy() {
					ud.Elem.Get("dataset").Set("originalDisplay", current)
				} else {
					ud.Elem.Get("dataset").Set("originalDisplay", "block")
				}
				style.Set("display", "none")
			}
			return 0
		},

		// ---------------------------
		// Clases CSS
		// ---------------------------
		"addClass": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			class := L.CheckString(2)
			ud.Elem.Get("classList").Call("add", class)
			return 0
		},
		"removeClass": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			class := L.CheckString(2)
			ud.Elem.Get("classList").Call("remove", class)
			return 0
		},
		"hasClass": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			class := L.CheckString(2)
			L.Push(lua.LBool(ud.Elem.Get("classList").Call("contains", class).Bool()))
			return 1
		},
		"toggleClass": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			class := L.CheckString(2)
			ud.Elem.Get("classList").Call("toggle", class)
			return 0
		},

		// ---------------------------
		// Atributos y propiedades
		// ---------------------------
		"setAttr": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			key := L.CheckString(2)
			val := L.CheckString(3)
			ud.Elem.Set(key, val)
			return 0
		},
		"getAttr": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			key := L.CheckString(2)
			L.Push(lua.LString(ud.Elem.Get(key).String()))
			return 1
		},
		"hasAttr": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			key := L.CheckString(2)
			L.Push(lua.LBool(!ud.Elem.Get(key).IsUndefined()))
			return 1
		},
		"removeAttr": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			key := L.CheckString(2)
			ud.Elem.Call("removeAttribute", key)
			return 0
		},
		"setProps": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			props := L.CheckTable(2)
			props.ForEach(func(k, v lua.LValue) {
				ud.Elem.Set(k.String(), v.String())
			})
			return 0
		},

		// ---------------------------
		// Estilo y contenido
		// ---------------------------
		"setStyle": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			styles := L.CheckTable(2)
			styles.ForEach(func(k, v lua.LValue) {
				ud.Elem.Get("style").Set(k.String(), v.String())
			})
			return 0
		},
		"setText": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			text := L.CheckString(2)
			ud.Elem.Set("textContent", text)
			return 0
		},

		// ---------------------------
		// Eventos
		// ---------------------------
		"on": func(L *lua.LState) int {
			ud := checkDOM(L, 1)
			event := L.CheckString(2)
			callback := L.CheckFunction(3)

			cb := js.FuncOf(func(this js.Value, args []js.Value) any {
				evt := L.NewTable()

				targetUD := L.NewUserData()
				targetUD.Value = &DOMUserData{Elem: args[0].Get("target")}
				evt.RawSetString("target", targetUD)

				if args[0].Get("value").Truthy() {
					evt.RawSetString("value", lua.LString(args[0].Get("value").String()))
				}
				if args[0].Get("checked").Truthy() {
					evt.RawSetString("checked", lua.LBool(args[0].Get("checked").Bool()))
				}

				evt.RawSetString("preventDefault", L.NewFunction(func(L *lua.LState) int {
					args[0].Call("preventDefault")
					return 0
				}))
				evt.RawSetString("stopPropagation", L.NewFunction(func(L *lua.LState) int {
					args[0].Call("stopPropagation")
					return 0
				}))

				L.CallByParam(lua.P{Fn: callback, NRet: 0, Protect: true}, evt)
				return nil
			})

			ud.Elem.Call("addEventListener", event, cb)
			return 0
		},
	}))
	return mt
}

func checkDOM(L *lua.LState, idx int) *DOMUserData {
	ud := L.CheckUserData(idx)
	if v, ok := ud.Value.(*DOMUserData); ok {
		return v
	}
	L.ArgError(idx, "DOM element expected")
	return nil
}

// ---------------------------
// Funciones públicas del módulo
// ---------------------------
func newElement(tag string) lua.LGFunction {
	return func(L *lua.LState) int {
		props := L.OptTable(1, L.NewTable())
		var child lua.LValue
		if L.GetTop() >= 2 {
			child = L.Get(2)
		} else {
			child = lua.LNil
		}

		elem := js.Global().Get("document").Call("createElement", tag)
		props.ForEach(func(k, v lua.LValue) {
			elem.Set(k.String(), v.String())
		})

		switch child.Type() {
		case lua.LTString:
			elem.Set("textContent", child.String())
		case lua.LTUserData:
			if c, ok := child.(*lua.LUserData); ok {
				if d, ok := c.Value.(*DOMUserData); ok {
					elem.Call("appendChild", d.Elem)
				}
			}
		case lua.LTTable:
			child.(*lua.LTable).ForEach(func(_, v lua.LValue) {
				if c, ok := v.(*lua.LUserData); ok {
					if d, ok := c.Value.(*DOMUserData); ok {
						elem.Call("appendChild", d.Elem)
					}
				}
			})
		}

		ud := L.NewUserData()
		ud.Value = &DOMUserData{Elem: elem}
		L.SetMetatable(ud, L.GetTypeMetatable(domTypeName))
		L.Push(ud)
		return 1
	}
}

func domRoot(L *lua.LState) int {
	elem := js.Global().Get("document").Get("body")
	ud := L.NewUserData()
	ud.Value = &DOMUserData{Elem: elem}
	L.SetMetatable(ud, L.GetTypeMetatable(domTypeName))
	L.Push(ud)
	return 1
}

// ---------------------------
// Loader del módulo
// ---------------------------
func Loader(L *lua.LState) *lua.LTable {
	registerMethods(L)
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"div":    newElement("div"),
		"span":   newElement("span"),
		"button": newElement("button"),
		"input":  newElement("input"),
		"root":   domRoot,
		"create": func(L *lua.LState) int {
			tag := L.CheckString(1)
			props := L.OptTable(2, L.NewTable())
			var child lua.LValue
			if L.GetTop() >= 3 {
				child = L.Get(3)
			} else {
				child = lua.LNil
			}
			elem := js.Global().Get("document").Call("createElement", tag)
			props.ForEach(func(k, v lua.LValue) {
				elem.Set(k.String(), v.String())
			})
			switch child.Type() {
			case lua.LTString:
				elem.Set("textContent", child.String())
			case lua.LTUserData:
				if c, ok := child.(*lua.LUserData); ok {
					if d, ok := c.Value.(*DOMUserData); ok {
						elem.Call("appendChild", d.Elem)
					}
				}
			case lua.LTTable:
				child.(*lua.LTable).ForEach(func(_, v lua.LValue) {
					if c, ok := v.(*lua.LUserData); ok {
						if d, ok := c.Value.(*DOMUserData); ok {
							elem.Call("appendChild", d.Elem)
						}
					}
				})
			}
			ud := L.NewUserData()
			ud.Value = &DOMUserData{Elem: elem}
			L.SetMetatable(ud, L.GetTypeMetatable(domTypeName))
			L.Push(ud)
			return 1
		},
	})
	return mod
}
