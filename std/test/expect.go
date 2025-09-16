//go:build !js

package test

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func Expect(L *lua.LState) int {
	val := L.CheckAny(1)
	ud := L.NewUserData()
	ud.Value = val
	L.SetMetatable(ud, L.GetTypeMetatable("expect"))
	L.Push(ud)
	return 1
}

func expectToEqual(L *lua.LState) int {
	ud := L.CheckUserData(1)
	expected := L.CheckAny(2)

	if ud.Value == expected {
		// Success: return nil (no error)
		L.Push(lua.LNil)
		return 1
	}

	// Error: return error string
	msg := fmt.Sprintf("expected %v, got %v", expected, ud.Value)
	L.Push(lua.LString(msg))
	return 1
}

func registerExpectType(L *lua.LState) {
	mt := L.NewTypeMetatable("expect")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"to_equal": expectToEqual,
		// Puedes añadir más métodos: to_not_equal, to_be_true, etc
	}))
}
