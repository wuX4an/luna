//go:build js && wasm
// +build js,wasm

package animate

import (
	"fmt"
	"luna/std/web/dom"
	"regexp"
	"strconv"
	"strings"
	"syscall/js"

	lua "github.com/yuin/gopher-lua"
)

// Registro global de animaciones activas por elemento

var (
	animIDCounter     int64 = 0
	elementAnimations       = map[string][]string{} // clave = elemento único
)

func getElementID(elem js.Value) string {
	dataset := elem.Get("dataset")
	id := dataset.Get("__luna_anim_id")
	if id.Truthy() {
		return id.String()
	}
	// Asignar nuevo ID
	animIDCounter++
	newID := strconv.FormatInt(animIDCounter, 10)
	dataset.Set("__luna_anim_id", newID)
	return newID
}

// Loader del módulo
func Loader(L *lua.LState) *lua.LTable {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"animate": animateFunc,
		"clear":   clearFunc,
	})
	return mod
}

// --- Función para animar ---
func animateFunc(L *lua.LState) int {
	ud := L.CheckUserData(1)
	opts := L.CheckTable(2)

	elemUD, ok := ud.Value.(*dom.DOMUserData)
	if !ok {
		L.ArgError(1, "DOM element expected")
		return 0
	}
	elem := elemUD.Elem

	// --- Opciones ---
	duration := getInt(opts, "duration", 1000)
	delay := getInt(opts, "delay", 0)
	easing := getString(opts, "easing", "ease")
	loop := getBool(opts, "loop", false)
	fill := getString(opts, "fill", "forwards")
	keep := getBool(opts, "keep", false)
	trigger := getString(opts, "trigger", "auto")

	xPtr := getFloatPtr(opts, "x")
	yPtr := getFloatPtr(opts, "y")
	sPtr := getFloatPtr(opts, "scale")
	rPtr := getFloatPtr(opts, "rotate")
	oPtr := getFloatPtr(opts, "opacity")

	if xPtr == nil && yPtr == nil && sPtr == nil && rPtr == nil && oPtr == nil {
		L.Push(lua.LNil)
		return 1
	}

	// --- Construir nombres únicos ---
	perf := js.Global().Get("performance").Call("now").Int()
	animName := fmt.Sprintf("luna_anim_%d", perf)
	className := "luna-anim-" + animName

	// --- Transform ---
	fromParts, toParts := []string{}, []string{}
	if xPtr != nil {
		fromParts = append(fromParts, "translateX(0px)")
		toParts = append(toParts, fmt.Sprintf("translateX(%vpx)", *xPtr))
	}
	if yPtr != nil {
		fromParts = append(fromParts, "translateY(0px)")
		toParts = append(toParts, fmt.Sprintf("translateY(%vpx)", *yPtr))
	}
	if rPtr != nil {
		fromParts = append(fromParts, "rotate(0deg)")
		toParts = append(toParts, fmt.Sprintf("rotate(%vdeg)", *rPtr))
	}
	if sPtr != nil {
		fromParts = append(fromParts, "scale(1)")
		toParts = append(toParts, fmt.Sprintf("scale(%v)", *sPtr))
	}
	fromTransform, toTransform := "none", "none"
	if len(fromParts) > 0 {
		fromTransform = joinTransformParts(fromParts)
		toTransform = joinTransformParts(toParts)
	}
	fromOpacity, toOpacity := "", ""
	if oPtr != nil {
		fromOpacity = "opacity: 1;"
		toOpacity = fmt.Sprintf("opacity: %v;", *oPtr)
	}
	loopStr := "1"
	if loop {
		loopStr = "infinite"
	}

	// --- CSS string ---
	selector := "." + className
	if trigger == "hover" {
		selector += ":hover"
	}

	css := fmt.Sprintf(`@keyframes %s { from { transform: %s; %s } to { transform: %s; %s } } %s { animation-name: %s; animation-duration: %dms; animation-timing-function: %s; animation-iteration-count: %s; animation-fill-mode: %s; animation-delay: %dms; }`,
		animName,
		fromTransform, fromOpacity,
		toTransform, toOpacity,
		selector,
		animName, duration, easing, loopStr, fill, delay)

	// --- Style global único ---
	doc := js.Global().Get("document")
	head := doc.Get("head")
	style := doc.Call("getElementById", "luna-anim-style")
	if !style.Truthy() {
		style = doc.Call("createElement", "style")
		style.Set("id", "luna-anim-style")
		style.Set("type", "text/css")
		head.Call("appendChild", style)
	}
	style.Set("textContent", style.Get("textContent").String()+"\n"+css)

	// Forzar reflow
	_ = elem.Get("offsetWidth").Int()

	// --- Añadir clase ---
	elem.Get("classList").Call("add", className)

	// Registrar animación activa
	key := getElementID(elem)
	elementAnimations[key] = append(elementAnimations[key], className)

	// --- Callback ---
	if trigger == "auto" {
		onComplete := opts.RawGetString("onComplete")
		hasCallback := onComplete != lua.LNil && (onComplete.Type() == lua.LTFunction)

		var cb js.Func
		cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if !keep {
				// Borrar animación temporal
				clearElementAnimation(elem, animName, className, style)
			}

			if hasCallback {
				targetUD := L.NewUserData()
				targetUD.Value = elemUD
				L.CallByParam(lua.P{Fn: onComplete, NRet: 0, Protect: true}, targetUD)
			}

			elem.Call("removeEventListener", "animationend", cb)
			cb.Release()
			return nil
		})
		elem.Call("addEventListener", "animationend", cb)
	}

	L.Push(lua.LNil)
	return 1
}

// --- Función para limpiar todas las animaciones de un elemento ---
func clearFunc(L *lua.LState) int {
	ud := L.CheckUserData(1)
	elemUD, ok := ud.Value.(*dom.DOMUserData)
	if !ok {
		L.ArgError(1, "DOM element expected")
		return 0
	}
	elem := elemUD.Elem

	style := js.Global().Get("document").Call("getElementById", "luna-anim-style")
	if !style.Truthy() {
		L.Push(lua.LNil)
		return 1
	}

	key := getElementID(elem)
	anims := elementAnimations[key]
	for _, className := range anims {
		animName := strings.TrimPrefix(className, "luna-anim-")
		clearElementAnimation(elem, animName, className, style)
	}

	elementAnimations[key] = []string{}
	L.Push(lua.LNil)
	return 1
}

// --- Helper para borrar clase y keyframe de un estilo ---
func clearElementAnimation(elem js.Value, animName, className string, style js.Value) {
	elem.Get("classList").Call("remove", className)
	content := style.Get("textContent").String()
	re := regexp.MustCompile(fmt.Sprintf(`(?s)@keyframes %s {.*?}\s*\.%s {.*?}`, animName, className))
	style.Set("textContent", re.ReplaceAllString(content, ""))
}

// --- Helpers ---
func getInt(tbl *lua.LTable, key string, def int) int {
	val := tbl.RawGetString(key)
	if num, ok := val.(lua.LNumber); ok {
		return int(num)
	}
	return def
}

func getString(tbl *lua.LTable, key, def string) string {
	val := tbl.RawGetString(key)
	if str, ok := val.(lua.LString); ok {
		return str.String()
	}
	return def
}

func getBool(tbl *lua.LTable, key string, def bool) bool {
	val := tbl.RawGetString(key)
	if val.Type() == lua.LTBool {
		return lua.LVAsBool(val)
	}
	return def
}

func getFloatPtr(tbl *lua.LTable, key string) *float64 {
	val := tbl.RawGetString(key)
	if num, ok := val.(lua.LNumber); ok {
		f := float64(num)
		return &f
	}
	return nil
}

func joinTransformParts(parts []string) string {
	return strings.Join(parts, " ")
}
