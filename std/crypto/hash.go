package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

	lua "github.com/yuin/gopher-lua"
)

// -----------------------------
// crypto.hash(algorithm, data)
func Hash(L *lua.LState) int {
	algorithm := L.CheckString(1)
	data := L.CheckString(2)

	var hashBytes []byte
	switch algorithm {
	case "md5":
		h := md5.Sum([]byte(data))
		hashBytes = h[:]
	case "sha1":
		h := sha1.Sum([]byte(data))
		hashBytes = h[:]
	case "sha256":
		h := sha256.Sum256([]byte(data))
		hashBytes = h[:]
	default:
		L.RaiseError("unsupported hash algorithm: %s", algorithm)
		return 0
	}

	L.Push(lua.LString(hex.EncodeToString(hashBytes)))
	return 1
}
