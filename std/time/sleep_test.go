package time

import (
	"testing"
	"time"

	lua "github.com/yuin/gopher-lua"
)

func TestSleep(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	start := time.Now()

	L.Push(lua.LNumber(100)) // 100 ms
	if err := L.CallByParam(lua.P{
		Fn:      L.NewFunction(sleep),
		NRet:    0,
		Protect: true,
	}, L.Get(-1)); err != nil {
		t.Fatal(err)
	}

	elapsed := time.Since(start)
	if elapsed < 100*time.Millisecond {
		t.Errorf("Expected sleep to last at least 100ms, got %v", elapsed)
	}
}
