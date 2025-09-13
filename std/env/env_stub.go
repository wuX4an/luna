//go:build js && wasm

package env

import (
	"luna/src/stub"

	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	return stub.NewStub("std:env")(L)
}
