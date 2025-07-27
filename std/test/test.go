package test

import (
	lua "github.com/yuin/gopher-lua"
)

type Test struct {
	// aquí puedes guardar estado si quieres (contador, reportes, etc)
}

func Loader(L *lua.LState) int {
	registerExpectType(L)
	registerTestType(L)

	ud := L.NewUserData()
	ud.Value = &Test{}
	L.SetMetatable(ud, L.GetTypeMetatable("test"))

	L.Push(ud)
	return 1
}

func registerTestType(L *lua.LState) {
	mod := L.NewTypeMetatable("test")
	L.SetField(mod, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"run":      Run,
		"expect":   Expect,
		"describe": Describe,
		"each":     Each,
	}))

	L.SetField(mod, "__call", L.NewFunction(RunOnCall)) // __call para llamar a test(...)
}
