//go:build !js

package test

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func runTestFunc(desc string, fn *lua.LFunction, L *lua.LState) error {
	// chequear variable de entorno
	if os.Getenv("LUNA_RUN_TESTS") != "1" {
		// solo registrar en pendingTests y no ejecutar
		tbl := L.NewTable()
		L.SetField(tbl, "name", lua.LString(desc))
		L.SetField(tbl, "fn", fn)

		return nil
	}

	// ejecución normal
	fmt.Printf("\x1b[36mRunning test:\x1b[0m %s\n", desc)
	err := L.CallByParam(lua.P{Fn: fn, NRet: 1, Protect: true})
	if err != nil {
		fmt.Printf("\x1b[31m✖ Assertion failed in\x1b[0m %s:\n", desc)
		fmt.Printf("\x1b[90m → %s\x1b[0m\n", err.Error())
		return err
	}
	ret := L.Get(-1)
	L.Pop(1)
	if str, ok := ret.(lua.LString); ok && str != "" {
		fmt.Printf("\x1b[31m✖ Assertion failed in\x1b[0m %s:\n", desc)
		fmt.Printf("\x1b[90m → %s\x1b[0m\n", string(str))
		return fmt.Errorf(string(str))
	}
	fmt.Printf("\x1b[32m✔ Test passed:\x1b[0m %s\n", desc)
	return nil
}

func Run(L *lua.LState) int {
	desc := L.CheckString(1)
	fn := L.CheckFunction(2)
	runTestFunc(desc, fn, L)
	return 0
}

func RunOnCall(L *lua.LState) int {
	desc := L.CheckString(2)
	fn := L.CheckFunction(3)
	runTestFunc(desc, fn, L)
	return 0
}
