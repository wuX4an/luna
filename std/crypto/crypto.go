package crypto

import lua "github.com/yuin/gopher-lua"

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"hash":         Hash,
		"hmac":         Hmac,
		"random_bytes": RandomBytes,
		"uuid":         Uuid,
	})
	L.Push(mod)
	return 1
}
