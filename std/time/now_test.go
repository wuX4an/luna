package time

import (
	"testing"

	lua "github.com/yuin/gopher-lua"
)

func TestNow(t *testing.T) {
	L := lua.NewState()
	defer L.Close()
	err := L.CallByParam(lua.P{
		Fn:      L.NewFunction(now),
		NRet:    1,
		Protect: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	ret := L.Get(-1)
	if _, ok := ret.(lua.LNumber); !ok {
		t.Errorf("Expected number, got %v", ret.Type())
	}
}
