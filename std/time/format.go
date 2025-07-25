package time

import (
	"fmt"
	"strings"
	"time"

	lua "github.com/yuin/gopher-lua"
)

// format(fmt, timestamp) → string
func format(L *lua.LState) int {
	formatStr := L.CheckString(1)
	timestamp := L.CheckNumber(2)

	t := time.Unix(int64(timestamp), 0).Local()

	// Primero, traducir el formato strftime a Go con marcadores temporales
	goFmt := goTimeFormat(formatStr)

	formatted := t.Format(goFmt)

	// Reemplazar manualmente los marcadores que Go no maneja:
	formatted = strings.ReplaceAll(formatted, "__DAY_FULL__", t.Weekday().String())
	formatted = strings.ReplaceAll(formatted, "__DAY_SHORT__", t.Weekday().String()[:3])

	monthFull := t.Month().String()
	formatted = strings.ReplaceAll(formatted, "__MONTH_FULL__", monthFull)
	formatted = strings.ReplaceAll(formatted, "__MONTH_SHORT__", monthFull[:3])

	formatted = strings.ReplaceAll(formatted, "__TZ__", t.Location().String())

	// Ahora reemplazamos __OFFSET__ con el offset horario local en formato ±HHMM
	_, offsetSeconds := t.Zone()
	sign := "+"
	if offsetSeconds < 0 {
		sign = "-"
		offsetSeconds = -offsetSeconds
	}
	hours := offsetSeconds / 3600
	minutes := (offsetSeconds % 3600) / 60
	offset := fmt.Sprintf("%s%02d%02d", sign, hours, minutes)

	formatted = strings.ReplaceAll(formatted, "__OFFSET__", offset)

	L.Push(lua.LString(formatted))
	return 1
}

func goTimeFormat(fmt string) string {
	// Estos sí son directamente compatibles con time.Format
	replacements := map[string]string{
		"%Y": "2006",
		"%y": "06",
		"%m": "01",
		"%d": "02",
		"%H": "15",
		"%I": "03",
		"%M": "04",
		"%S": "05",
		"%p": "PM",
		"%z": "__OFFSET__",      // <-- cambiar aquí
		"%Z": "__TZ__",          // manual
		"%A": "__DAY_FULL__",    // manual
		"%a": "__DAY_SHORT__",   // manual
		"%B": "__MONTH_FULL__",  // manual
		"%b": "__MONTH_SHORT__", // manual
		"%F": "2006-01-02",
		"%T": "15:04:05",
		"%r": "03:04:05 PM",
		"%R": "15:04",
	}
	for k, v := range replacements {
		fmt = strings.ReplaceAll(fmt, k, v)
	}
	return fmt
}
