package base64

import lua "github.com/yuin/gopher-lua"

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"encode": Encode,
		"decode": Decode,
	})
	L.Push(mod)
	return 1
}
