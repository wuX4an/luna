//go:build js && wasm

package sqlite

import (
	"luna/src/stub"

	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	return stub.NewStub("std:sqlite")(L)
}
