package time

import (
	"testing"
	"time"

	lua "github.com/yuin/gopher-lua"
)

func TestFormat(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	// Creamos un tiempo local para el 24 de julio 2024 a las 18:29:32
	refTime := time.Date(2024, time.July, 24, 18, 29, 32, 0, time.Local)
	timestamp := lua.LNumber(refTime.Unix())

	tests := []struct {
		fmt      string
		expected string
	}{
		{"%Y-%m-%d", "2024-07-24"},
		{"%A", refTime.Weekday().String()},
		{"%a", refTime.Weekday().String()[:3]},
		{"%B", refTime.Month().String()},
		{"%b", refTime.Month().String()[:3]},
		{"%H:%M:%S", "18:29:32"},
		{"%I:%M %p", "06:29 PM"},
		{"%Y/%m/%d %H:%M", "2024/07/24 18:29"},
	}

	for _, test := range tests {
		L.Push(lua.LString(test.fmt))
		L.Push(timestamp)

		err := L.CallByParam(lua.P{
			Fn:      L.NewFunction(format),
			NRet:    1,
			Protect: true,
		}, L.Get(-2), L.Get(-1))
		if err != nil {
			t.Fatalf("format(%q) returned error: %v", test.fmt, err)
		}

		result := L.Get(-1)
		str, ok := result.(lua.LString)
		if !ok {
			t.Fatalf("format(%q) expected string result, got %s", test.fmt, result.Type())
		}

		if str != lua.LString(test.expected) {
			t.Errorf("format(%q) = %q; want %q", test.fmt, str, test.expected)
		}

		L.Pop(1) // limpiar resultado para la siguiente iteración
	}
}
