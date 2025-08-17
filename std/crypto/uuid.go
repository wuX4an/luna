package crypto

import (
	"github.com/google/uuid"
	lua "github.com/yuin/gopher-lua"
)

// -----------------------------
// crypto.uuid()
func Uuid(L *lua.LState) int {
	id := uuid.New()
	L.Push(lua.LString(id.String()))
	return 1
}
