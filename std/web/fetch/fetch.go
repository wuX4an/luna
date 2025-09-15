//go:build js && wasm

package fetch

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// ---------------------------
// Loader del módulo
// ---------------------------
func Loader(L *lua.LState) *lua.LTable {
	mod := L.NewTable()
	mod.RawSetString("fetch", L.NewFunction(Fetch))
	mod.RawSetString("get", L.NewFunction(Get))
	mod.RawSetString("post", L.NewFunction(Post))
	return mod
}

// ---------------------------
// Funciones
// ---------------------------

// fetch(url, opts, callback)
func Fetch(L *lua.LState) int {
	url := L.CheckString(1)

	// Opciones
	var opts js.Value
	if L.GetTop() >= 2 && L.Get(2).Type() == lua.LTTable {
		table := L.CheckTable(2)
		opts = js.Global().Get("Object").New()

		// method
		if m := table.RawGetString("method"); m != lua.LNil {
			opts.Set("method", lua.LVAsString(m))
		}
		// headers
		if h := table.RawGetString("headers"); h != lua.LNil {
			headers := js.Global().Get("Object").New()
			h.(*lua.LTable).ForEach(func(k, v lua.LValue) {
				headers.Set(lua.LVAsString(k), lua.LVAsString(v))
			})
			opts.Set("headers", headers)
		}
		// body
		if b := table.RawGetString("body"); b != lua.LNil {
			opts.Set("body", lua.LVAsString(b))
		}
	} else {
		opts = js.Undefined()
	}

	// Callback Lua
	callback := L.CheckFunction(L.GetTop())

	// Ejecutar fetch
	promise := js.Global().Call("fetch", url, opts)

	var thenFunc js.Func
	thenFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		resp := args[0]

		status := resp.Get("status").Int()
		ok := resp.Get("ok").Bool()

		var textFunc js.Func
		textFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
			body := args[0].String()

			result := L.NewTable()
			result.RawSetString("status", lua.LNumber(status))
			result.RawSetString("ok", lua.LBool(ok))
			result.RawSetString("body", lua.LString(body))

			L.CallByParam(lua.P{Fn: callback, NRet: 0, Protect: true}, result)

			textFunc.Release() // liberar función
			return nil
		})

		resp.Call("text").Call("then", textFunc)
		thenFunc.Release() // liberar función
		return nil
	})

	promise.Call("then", thenFunc)
	return 0
}

// fetch.get(url)
func Get(L *lua.LState) int {
	url := L.CheckString(1)
	// devuelve una función Lua que recibe un callback
	fn := L.NewFunction(func(L *lua.LState) int {
		callback := L.CheckFunction(1)

		// ejecutar fetch con JS
		promise := js.Global().Call("fetch", url)
		promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
			resp := args[0]
			textPromise := resp.Call("text")
			textPromise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
				result := L.NewTable()
				result.RawSetString("status", lua.LNumber(resp.Get("status").Int()))
				result.RawSetString("ok", lua.LBool(resp.Get("ok").Bool()))
				result.RawSetString("body", lua.LString(args[0].String()))

				// llamar callback
				L.CallByParam(lua.P{
					Fn:      callback,
					NRet:    0,
					Protect: true,
				}, result)
				return nil
			}))
			return nil
		}))

		return 0
	})

	L.Push(fn)
	return 1
}

// fetch.post(url, opts)
func Post(L *lua.LState) int {
	url := L.CheckString(1)
	opts := L.CheckTable(2)

	// Devolver una función Lua que recibe un callback
	fn := L.NewFunction(func(L *lua.LState) int {
		callback := L.CheckFunction(1)

		// Crear JS options
		jsOpts := js.Global().Get("Object").New()
		jsOpts.Set("method", "POST")

		if h := opts.RawGetString("headers"); h != lua.LNil {
			headers := js.Global().Get("Object").New()
			h.(*lua.LTable).ForEach(func(k, v lua.LValue) {
				headers.Set(lua.LVAsString(k), lua.LVAsString(v))
			})
			jsOpts.Set("headers", headers)
		}
		if b := opts.RawGetString("body"); b != lua.LNil {
			jsOpts.Set("body", lua.LVAsString(b))
		}

		// Ejecutar fetch JS
		promise := js.Global().Call("fetch", url, jsOpts)

		var thenFunc js.Func
		thenFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
			resp := args[0]

			status := resp.Get("status").Int()
			ok := resp.Get("ok").Bool()

			var textFunc js.Func
			textFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
				body := args[0].String()

				result := L.NewTable()
				result.RawSetString("status", lua.LNumber(status))
				result.RawSetString("ok", lua.LBool(ok))
				result.RawSetString("body", lua.LString(body))

				L.CallByParam(lua.P{Fn: callback, NRet: 0, Protect: true}, result)

				textFunc.Release()
				return nil
			})

			resp.Call("text").Call("then", textFunc)
			thenFunc.Release()
			return nil
		})

		promise.Call("then", thenFunc)
		return 0
	})

	L.Push(fn)
	return 1
}
