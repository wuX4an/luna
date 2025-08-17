package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"

	lua "github.com/yuin/gopher-lua"
)

// -----------------------------
// crypto.hmac(algorithm, key, data)
func Hmac(L *lua.LState) int {
	algorithm := L.CheckString(1)
	key := []byte(L.CheckString(2))
	data := []byte(L.CheckString(3))

	var mac hash.Hash
	switch algorithm {
	case "sha1":
		mac = hmac.New(sha1.New, key)
	case "sha256":
		mac = hmac.New(sha256.New, key)
	default:
		L.RaiseError("unsupported HMAC algorithm: %s", algorithm)
		return 0
	}

	mac.Write(data)
	L.Push(lua.LString(hex.EncodeToString(mac.Sum(nil))))
	return 1
}
