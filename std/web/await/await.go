//go:build js && wasm
// +build js,wasm

package await

import (
	lua "github.com/yuin/gopher-lua"
)

// Loader del módulo await: devuelve directamente la función
func Loader(L *lua.LState) *lua.LFunction {
	return L.NewFunction(Await)
}

// await(promiseFn) -> suspende la coroutine hasta que la función llame al callback
func Await(L *lua.LState) int {
	// promiseFn es una función Lua que recibirá un callback
	promiseFn := L.CheckFunction(1)

	// canal para recibir el resultado
	resultChan := make(chan lua.LValue, 1)

	// callback Lua que la función promiseFn debe llamar
	callback := L.NewFunction(func(L *lua.LState) int {
		res := L.Get(1)
		resultChan <- res
		return 0
	})

	// llamar a promiseFn pasando el callback
	L.Push(promiseFn)
	L.Push(callback)
	L.Call(1, 0) // 1 argumento: callback

	// suspender hasta recibir el resultado
	res := <-resultChan

	// devolver el resultado a Lua
	L.Push(res)
	return 1
}
