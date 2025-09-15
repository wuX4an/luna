//go:build js && wasm

package audio

import (
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

type LuaAudio struct {
	obj js.Value
}

// ---------------------------
// Loader del módulo
// ---------------------------
func Loader(L *lua.LState) *lua.LTable {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"new": NewAudio,
	})
	return mod
}

// ---------------------------
// Funciones del módulo
// ---------------------------
func NewAudio(L *lua.LState) int {
	src := L.CheckString(1)
	audioObj := js.Global().Get("Audio").New(src)
	luaAudio := &LuaAudio{obj: audioObj}

	// Crear userdata
	ud := L.NewUserData()
	ud.Value = luaAudio

	// Crear metatable global
	mt := L.NewTypeMetatable("audio_mt")
	// Tabla de métodos
	methods := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"play":           audioPlay,
		"pause":          audioPause,
		"stop":           audioStop,
		"setVolume":      audioSetVolume,
		"getVolume":      audioGetVolume,
		"setLoop":        audioSetLoop,
		"setRate":        audioSetRate,
		"setPosition":    audioSetPosition,
		"getPosition":    audioGetPosition,
		"getMaxPosition": audioGetMaxPosition,
		"getMinPosition": audioGetMinPosition,
		"getDuration":    audioGetDuration,
	})
	mt.RawSetString("__index", methods)

	L.SetMetatable(ud, mt)
	L.Push(ud)
	return 1
}

// ---------------------------
// Helpers
// ---------------------------
func checkAudio(L *lua.LState) *LuaAudio {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*LuaAudio); ok {
		return v
	}
	L.ArgError(1, "Audio expected")
	return nil
}

// ---------------------------
// Métodos
// ---------------------------
func audioPlay(L *lua.LState) int {
	a := checkAudio(L)
	a.obj.Call("play")
	return 0
}

func audioPause(L *lua.LState) int {
	a := checkAudio(L)
	a.obj.Call("pause")
	return 0
}

func audioStop(L *lua.LState) int {
	a := checkAudio(L)
	a.obj.Call("pause")
	a.obj.Set("currentTime", 0)
	return 0
}

func audioSetVolume(L *lua.LState) int {
	a := checkAudio(L)
	vol := L.CheckNumber(2)
	a.obj.Set("volume", float64(vol))
	return 0
}

func audioGetVolume(L *lua.LState) int {
	a := checkAudio(L)
	vol := a.obj.Get("volume").Float()
	L.Push(lua.LNumber(vol))
	return 1
}

func audioSetLoop(L *lua.LState) int {
	a := checkAudio(L)
	loop := L.CheckBool(2)
	a.obj.Set("loop", loop)
	return 0
}

func audioSetRate(L *lua.LState) int {
	a := checkAudio(L)
	rate := L.CheckNumber(2)
	a.obj.Set("playbackRate", float64(rate))
	return 0
}

func audioSetPosition(L *lua.LState) int {
	a := checkAudio(L)
	pos := L.CheckNumber(2)
	a.obj.Set("currentTime", float64(pos))
	return 0
}

func audioGetPosition(L *lua.LState) int {
	a := checkAudio(L)
	pos := a.obj.Get("currentTime").Float()
	L.Push(lua.LNumber(pos))
	return 1
}

func audioGetMaxPosition(L *lua.LState) int {
	a := checkAudio(L)
	dur := a.obj.Get("duration").Float()
	L.Push(lua.LNumber(dur))
	return 1
}

func audioGetMinPosition(L *lua.LState) int {
	// Siempre 0 en Audio de JS
	L.Push(lua.LNumber(0))
	return 1
}

func audioGetDuration(L *lua.LState) int {
	a := checkAudio(L)
	dur := a.obj.Get("duration").Float()
	L.Push(lua.LNumber(dur))
	return 1
}
