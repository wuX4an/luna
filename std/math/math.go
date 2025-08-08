package math

import lua "github.com/yuin/gopher-lua"

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"pow":  Pow,
		"sqrt": Sqrt,
		"log":  Log,
		"fact": Fact,
		// Trigonometry
		"cos": Cos,
		"sin": Sin,
		"tan": Tan,
	})

	L.SetField(mod, "pi", lua.LNumber(3.14159265358979323846))
	L.SetField(mod, "e", lua.LNumber(2.71828182845904523536))
	L.SetField(mod, "tau", lua.LNumber(6.28318530717958647692))
	L.SetField(mod, "phi", lua.LNumber(1.61803398874989484820))
	L.SetField(mod, "G", lua.LNumber(6.67430e-11)) // en m³·kg⁻¹·s⁻²

	L.Push(mod)
	return 1
}
