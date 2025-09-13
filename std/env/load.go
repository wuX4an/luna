//go:build !js

package env

import (
	"bufio"
	"os"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

var vars = map[string]string{}

func Load(L *lua.LState) int {
	nargs := L.GetTop()
	if nargs != 1 {
		L.RaiseError(
			"expected exactly 1 argument (string: filename), got %d",
			nargs,
		)
		return 0
	}

	filename := L.CheckString(1)
	file, err := os.Open(filename)
	if err != nil {
		L.RaiseError("failed to open file: %v", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		vars[key] = value
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		L.RaiseError("error reading file: %v", err)
		return 0
	}

	L.Push(lua.LTrue)
	return 1
}
