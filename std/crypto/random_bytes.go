package crypto

import (
	"crypto/rand"

	lua "github.com/yuin/gopher-lua"
)

// -----------------------------
// crypto.random_bytes(n)
func RandomBytes(L *lua.LState) int {
	n := L.CheckInt(1)
	if n <= 0 {
		L.RaiseError("n must be positive")
		return 0
	}

	buf := make([]byte, n)
	_, err := rand.Read(buf)
	if err != nil {
		L.RaiseError("failed to generate random bytes: %v", err)
		return 0
	}

	L.Push(lua.LString(buf))
	return 1
}
