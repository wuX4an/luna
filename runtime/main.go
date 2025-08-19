package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"luna/src/luavm"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

const (
	marker      = "LUNA_BUNDLE"
	markerSize  = len(marker)
	offsetBytes = 8
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Open(os.Args[0])
	if err != nil {
		return err
	}
	defer f.Close()

	offset, err := findTarOffset(f)
	if err != nil {
		return err
	}
	_, err = f.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	// --- Guardar todos los módulos en memoria ---
	modules := map[string]string{}
	var mainLuaContent string

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if strings.HasSuffix(hdr.Name, ".lua") {
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, tr); err != nil {
				return err
			}

			name := strings.TrimSuffix(filepath.Base(hdr.Name), ".lua")
			modules[name] = buf.String()

			if hdr.Name == "main.lua" {
				mainLuaContent = buf.String()
			}
		}
	}

	if mainLuaContent == "" {
		return errors.New("main.lua not found in bundle")
	}

	// --- Inicializar VM ---
	L := luavm.NewLuaVM()
	defer L.Close()

	originRequire := L.GetGlobal("require")
	// --- Sobrescribir require ---
	L.SetGlobal("require", L.NewFunction(func(L *lua.LState) int {
		modName := L.ToString(1)
		if code, ok := modules[modName]; ok {
			// Cargar desde bundle
			fn, err := L.LoadString(code)
			if err != nil {
				L.RaiseError("error compiling module %s: %v", modName, err)
				return 0
			}
			L.Push(fn)
			if err := L.PCall(0, 1, nil); err != nil {
				L.RaiseError("error running module %s: %v", modName, err)
				return 0
			}
			return 1
		}

		// Fallback: usar el require original
		L.Push(originRequire)
		L.Push(lua.LString(modName))
		if err := L.PCall(1, 1, nil); err != nil {
			L.RaiseError("fallback require failed for %s: %v", modName, err)
			return 0
		}
		return 1
	}))

	// --- Ejecutar main.lua ---
	if err := L.DoString(mainLuaContent); err != nil {
		return fmt.Errorf("error running main.lua: %w", err)
	}

	return nil
}

func findTarOffset(f *os.File) (int64, error) {
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}
	fileSize := stat.Size()
	if fileSize < int64(markerSize+offsetBytes) {
		return 0, errors.New("file too small to contain marker")
	}
	buf := make([]byte, markerSize+offsetBytes)
	_, err = f.ReadAt(buf, fileSize-int64(markerSize+offsetBytes))
	if err != nil {
		return 0, err
	}
	if string(buf[:markerSize]) != marker {
		return 0, errors.New("marker not found")
	}
	offset := int64(binary.LittleEndian.Uint64(buf[markerSize:]))
	return offset, nil
}
