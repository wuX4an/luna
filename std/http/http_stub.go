//go:build js && wasm

package http

import (
	"luna/src/stub"

	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	return stub.NewStub("std:http")(L)
}
