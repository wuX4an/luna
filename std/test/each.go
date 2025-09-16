//go:build !js

package test

import (
	lua "github.com/yuin/gopher-lua"
)

// each receives (cases table, fn function)
// Runs fn(case) for each case in the cases table,
// calling runTestFunc with a generated name and the Lua function
func Each(L *lua.LState) int {
	// First argument: table with cases
	cases := L.CheckTable(1)
	// Second argument: function to execute with each case
	fn := L.CheckFunction(2)

	// Iterate over the cases table
	cases.ForEach(func(_ lua.LValue, v lua.LValue) {
		// v is each case, which should be a table
		caseTbl, ok := v.(*lua.LTable)
		if !ok {
			// If not a table, raise an error
			L.RaiseError("test.each expects each case to be a table")
			return
		}

		// Get the name field of the case if it exists, otherwise use a default name
		nameLv := caseTbl.RawGetString("name")
		name := "param test"
		if str, ok := nameLv.(lua.LString); ok && str != "" {
			name = string(str)
		}

		// Create an anonymous function that calls fn(case)
		closure := L.NewFunction(func(L *lua.LState) int {
			L.Push(fn)
			L.Push(caseTbl)
			err := L.PCall(1, 1, nil)
			if err != nil {
				L.RaiseError("error running parametrized test: %v", err)
				return 0
			}
			return 1
		})

		runTestFunc(name, closure, L)
	})

	return 0
}
