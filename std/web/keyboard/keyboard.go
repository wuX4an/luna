//go:build js && wasm

package keyboard

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// ---------------------------
// Estado interno
// ---------------------------
var (
	pressedKeys        = map[string]bool{}
	disableAllDefaults = false
	disabledKeys       = map[string]bool{}
	keyListeners       []listener
	comboListeners     []comboListener
)

type listener struct {
	event    string
	callback *lua.LFunction
	once     bool
}

type comboListener struct {
	keys     []string
	callback *lua.LFunction
}

// ---------------------------
// Loader del módulo
// ---------------------------
func Loader(L *lua.LState) *lua.LTable {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"on":              on,
		"once":            once,
		"onCombo":         onCombo,
		"isPressed":       isPressed,
		"disableDefaults": disableDefaults,
	})
	setupEventListeners(L)
	return mod
}

// ---------------------------
// Funciones Lua
// ---------------------------

// on("keydown"/"keyup", callback)
func on(L *lua.LState) int {
	event := L.CheckString(1)
	callback := L.CheckFunction(2)
	keyListeners = append(keyListeners, listener{event: event, callback: callback})
	return 0
}

// once("keydown"/"keyup", callback)
func once(L *lua.LState) int {
	event := L.CheckString(1)
	callback := L.CheckFunction(2)
	keyListeners = append(keyListeners, listener{event: event, callback: callback, once: true})
	return 0
}

// onCombo({"Ctrl","S"}, callback)
func onCombo(L *lua.LState) int {
	keysTable := L.CheckTable(1)
	callback := L.CheckFunction(2)
	keys := []string{}
	keysTable.ForEach(func(_, v lua.LValue) {
		keys = append(keys, v.String())
	})
	comboListeners = append(comboListeners, comboListener{keys: keys, callback: callback})
	return 0
}

// isPressed("Shift")
func isPressed(L *lua.LState) int {
	key := L.CheckString(1)
	L.Push(lua.LBool(pressedKeys[key]))
	return 1
}

// disableDefaults() o disableDefaults({"Tab","F1"})
func disableDefaults(L *lua.LState) int {
	if L.GetTop() == 0 {
		disableAllDefaults = true
	} else {
		keysTable := L.CheckTable(1)
		keysTable.ForEach(func(_, v lua.LValue) {
			disabledKeys[v.String()] = true
		})
	}
	return 0
}

// ---------------------------
// Eventos JS
// ---------------------------
func setupEventListeners(L *lua.LState) {
	js.Global().Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handleKeyEvent(L, "keydown", args[0])
		return nil
	}))
	js.Global().Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handleKeyEvent(L, "keyup", args[0])
		return nil
	}))
}

func handleKeyEvent(L *lua.LState, eventType string, evt js.Value) {
	key := evt.Get("key").String()
	pressed := eventType == "keydown"

	// Actualizar estado
	pressedKeys[key] = pressed

	// Evitar default si corresponde
	if disableAllDefaults || disabledKeys[key] {
		evt.Call("preventDefault")
	}

	// Crear objeto Lua evt
	luaEvt := L.NewTable()
	luaEvt.RawSetString("key", lua.LString(key))
	luaEvt.RawSetString("code", lua.LString(evt.Get("code").String()))
	luaEvt.RawSetString("ctrlKey", lua.LBool(evt.Get("ctrlKey").Bool()))
	luaEvt.RawSetString("shiftKey", lua.LBool(evt.Get("shiftKey").Bool()))
	luaEvt.RawSetString("altKey", lua.LBool(evt.Get("altKey").Bool()))
	luaEvt.RawSetString("metaKey", lua.LBool(evt.Get("metaKey").Bool()))
	luaEvt.RawSetString("preventDefault", L.NewFunction(func(L2 *lua.LState) int {
		evt.Call("preventDefault")
		return 0
	}))
	luaEvt.RawSetString("stopPropagation", L.NewFunction(func(L2 *lua.LState) int {
		evt.Call("stopPropagation")
		return 0
	}))

	// Llamar listeners on/once
	newListeners := []listener{}
	for _, l := range keyListeners {
		if l.event == eventType {
			L.CallByParam(lua.P{
				Fn:      l.callback,
				NRet:    0,
				Protect: true,
			}, luaEvt)
			if !l.once {
				newListeners = append(newListeners, l)
			}
		} else {
			newListeners = append(newListeners, l)
		}
	}
	keyListeners = newListeners

	// Llamar combinaciones
	for _, cl := range comboListeners {
		if allPressed(cl.keys) {
			L.CallByParam(lua.P{
				Fn:      cl.callback,
				NRet:    0,
				Protect: true,
			}, luaEvt)
		}
	}
}

func allPressed(keys []string) bool {
	for _, k := range keys {
		if !pressedKeys[k] {
			return false
		}
	}
	return true
}
