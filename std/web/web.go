//go:build js && wasm
// +build js,wasm

package web

import (
	"luna/std/web/animate"
	"luna/std/web/audio"
	"luna/std/web/await"
	"luna/std/web/clipboard"
	"luna/std/web/console"
	"luna/std/web/cookie"
	"luna/std/web/dom"
	"luna/std/web/fetch"
	"luna/std/web/keyboard"
	"luna/std/web/notify"
	"luna/std/web/storage"
	"luna/std/web/timer"

	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	mod := L.NewTable()

	// Submódulos
	consoleMod := console.Loader(L)
	cookieMod := cookie.Loader(L)
	storageMod := storage.Loader(L)
	fetchMod := fetch.Loader(L)
	awaitMod := await.Loader(L)
	clipboardMod := clipboard.Loader(L)
	notifyMod := notify.Loader(L)
	timerMod := timer.Loader(L)
	domMod := dom.Loader(L)
	audioMod := audio.Loader(L)
	animateMod := animate.Loader(L)
	keyboardMod := keyboard.Loader(L)

	mod.RawSetString("console", consoleMod)
	mod.RawSetString("cookie", cookieMod)
	mod.RawSetString("storage", storageMod)
	mod.RawSetString("fetch", fetchMod)
	mod.RawSetString("await", awaitMod)
	mod.RawSetString("clipboard", clipboardMod)
	mod.RawSetString("notify", notifyMod)
	mod.RawSetString("timer", timerMod)
	mod.RawSetString("dom", domMod)
	mod.RawSetString("audio", audioMod)
	mod.RawSetString("animate", animateMod)
	mod.RawSetString("keyboard", keyboardMod)

	L.Push(mod)
	return 1
}
