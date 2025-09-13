//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"luna/src/luavm"
	"path/filepath"
	"strings"
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	// Comprovar si el runtime si inicializo correctamente
	js.Global().Get("console").Call("log", "Luna: ok")

	L := luavm.NewLuaVM()
	defer L.Close()

	// Cache de módulos locales en memoria
	modules := map[string]string{}

	// Función para fetch de módulos locales
	fetchLocalModule := func(name string) (string, error) {
		name = filepath.Join(strings.ReplaceAll(name, ".", "/"))

		if code, ok := modules[name]; ok {
			return code, nil
		}

		// Detectar Node/Bun: presencia de "process"
		isNode := js.Global().Get("process").Truthy() && js.Global().Get("require").Truthy()
		if isNode {
			fs := js.Global().Get("require").Invoke("fs")
			path := "_modules/" + name + ".lua"
			content := fs.Call("readFileSync", path, "utf8").String()
			modules[name] = content
			return content, nil
		}

		// Navegador
		url := "_modules/" + name + ".lua"
		js.Global().Get("console").Call("log", "Fetching module via fetch:", url)
		promise := js.Global().Call("fetch", url)
		then := make(chan js.Value)
		catch := make(chan js.Value)
		promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
			then <- args[0]
			return nil
		})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
			catch <- args[0]
			return nil
		}))
		select {
		case resp := <-then:
			textPromise := resp.Call("text")
			textThen := make(chan js.Value)
			textCatch := make(chan js.Value)
			textPromise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
				textThen <- args[0]
				return nil
			})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) any {
				textCatch <- args[0]
				return nil
			}))
			select {
			case t := <-textThen:
				code := t.String()
				modules[name] = code
				return code, nil
			case err := <-textCatch:
				return "", fmt.Errorf("failed reading module %s: %v", name, err)
			}
		case err := <-catch:
			return "", fmt.Errorf("failed fetching module %s: %v", name, err)
		}
	}

	// Guardar require original
	origRequire := L.GetGlobal("require")

	// Nuevo require
	L.SetGlobal("require", L.NewFunction(func(L *lua.LState) int {
		modName := L.ToString(1)

		// Si empieza con std:, usar require original tal cual
		if len(modName) >= 4 && modName[:4] == "std:" {
			L.Push(origRequire)
			L.Push(lua.LString(modName)) // NO cortamos el prefijo
			if err := L.PCall(1, 1, nil); err != nil {
				L.RaiseError("failed to require std module %s: %v", modName, err)
				return 0
			}
			return 1
		}

		// Si no es std, fetch desde _modules
		code, err := fetchLocalModule(modName)
		if err != nil {
			L.RaiseError("fetch failed for module %s: %v", modName, err)
			return 0
		}
		fn, err := L.LoadString(code)
		if err != nil {
			L.RaiseError("compile failed for module %s: %v", modName, err)
			return 0
		}
		L.Push(fn)
		if err := L.PCall(0, 1, nil); err != nil {
			L.RaiseError("execution failed for module %s: %v", modName, err)
			return 0
		}
		return 1
	}))

	// Cargar y ejecutar main.lua desde _modules
	mainCode, err := fetchLocalModule("main")
	if err != nil {
		js.Global().Get("console").Call("log", "Failed to fetch main.lua:", err.Error())
		return
	}

	if err := L.DoString(mainCode); err != nil {
		js.Global().Get("console").Call("log", "Failed to run main.lua:", err.Error())
		return
	}
}
