package cli

import (
	"fmt"
	"io/fs"
	"luna/src/config"
	"luna/src/luavm"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var testCmd = &cobra.Command{
	Use:     "test [FILE|DIR]",
	Short:   "Run tests",
	Example: "luna test main.lua\nluna test .",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Path default
		arg := "."
		if len(args) > 0 {
			arg = args[0]
		}

		info, err := os.Stat(arg)
		if err != nil {
			return fmt.Errorf("path does not exist: %s", arg)
		}

		var conf *config.LunaConfig
		var projectDir string
		var scriptPath string

		if info.IsDir() {
			configPath := filepath.Join(arg, "Luna.toml")
			conf, err = config.LoadConfig(configPath)
			if err != nil {
				return err
			}
			projectDir = filepath.Join(arg, conf.Build.Source)
			scriptPath = filepath.Join(projectDir, conf.Build.Entry)
		} else if filepath.Ext(arg) == ".lua" {
			projectDir = filepath.Dir(arg)
			scriptPath = arg
		} else {
			return fmt.Errorf("argument must be a .lua file or a project directory with Luna.toml")
		}

		// Inicializar VM
		L := luavm.NewLuaVM()
		defer L.Close()

		// Forzar ejecución de tests en este comando
		os.Setenv("LUNA_RUN_TESTS", "1")

		// Cargar módulo test
		if err := L.DoString(`test = require("std:test")`); err != nil {
			return fmt.Errorf("cannot load std:test: %w", err)
		}

		// Cargar todos los módulos Lua en memoria
		modules := map[string]string{}
		err = filepath.WalkDir(projectDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			if filepath.Ext(path) == ".lua" {
				rel, _ := filepath.Rel(projectDir, path)
				rel = filepath.ToSlash(rel)
				base := strings.TrimSuffix(rel, ".lua")
				modName := strings.ReplaceAll(base, "/", ":")
				data, _ := os.ReadFile(path)
				modules[modName] = string(data)
				modules[strings.ReplaceAll(base, "/", ".")] = string(data)
			}
			return nil
		})

		// Sobrescribir require para cargar desde memoria
		originRequire := L.GetGlobal("require")
		L.SetGlobal("require", L.NewFunction(func(L *lua.LState) int {
			modName := L.ToString(1)
			if code, ok := modules[modName]; ok {
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
			// fallback
			L.Push(originRequire)
			L.Push(lua.LString(modName))
			if err := L.PCall(1, 1, nil); err != nil {
				L.RaiseError("fallback require failed for %s: %v", modName, err)
				return 0
			}
			return 1
		}))

		// Ejecutar el script principal (registrará tests en __TESTS__)
		if err := L.DoFile(scriptPath); err != nil {
			return fmt.Errorf("error running script: %w", err)
		}

		// --- Ejecutar los tests de __TESTS__ si existen ---
		if err := L.DoString(`
if os.getenv("LUNA_RUN_TESTS") then
  for _, t in ipairs(__TESTS__ or {}) do t() end
end
`); err != nil {
			return fmt.Errorf("error executing pending tests: %w", err)
		}

		return nil
	},
}

func init() {
	Cmd.AddCommand(testCmd)
}
