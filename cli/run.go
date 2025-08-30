package cli

import (
	"fmt"
	"io/fs"
	"luna/src/config"
	"luna/src/luavm"
	"os"
	"path/filepath"
	"strings"

	handler "luna/src/error"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var runExamples = `  luna run main.lua
  luna run .`

var runCmd = &cobra.Command{
	Use:     "run [FILE|DIR]",
	Short:   "Run a Lua script or project",
	Example: runExamples,
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		// Default to current directory
		arg := "."
		if len(args) > 0 {
			arg = args[0]
		}

		info, err := os.Stat(arg)
		if err != nil {
			return fmt.Errorf("path does not exist: %s", arg)
		}

		var conf *config.LunaConfig
		var scriptPath string
		var projectDir string

		if info.IsDir() {
			// Load Luna.toml
			configPath := filepath.Join(arg, "Luna.toml")
			conf, err = config.LoadConfig(configPath)
			if err != nil {
				return err
			}

			scriptPath = filepath.Join(arg, conf.Build.Source, conf.Build.Entry)
			projectDir = filepath.Join(arg, conf.Build.Source)
		} else if filepath.Ext(arg) == ".lua" {
			// Single Lua file
			scriptPath = arg
			projectDir = filepath.Dir(arg)
		} else {
			return fmt.Errorf("argument must be a .lua file or a project directory with Luna.toml")
		}

		// Initialize Lua VM
		L := luavm.NewLuaVM()
		defer L.Close()

		// Load all Lua modules in memory
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

		// Override require to load from memory first
		originRequire := L.GetGlobal("require")
		L.SetGlobal("require", L.NewFunction(func(L *lua.LState) int {
			modName := L.ToString(1)
			if code, ok := modules[modName]; ok {
				modPath := filepath.Join(projectDir, strings.ReplaceAll(modName, ".", string(os.PathSeparator))+".lua")

				// Cargar y ejecutar el módulo SIN capturar errores con PCall

				fn, err := L.Load(strings.NewReader(code), modPath)
				if err != nil {
					// Solo para errores de sintaxis al cargar
					pretty := handler.PrettyLuaError(modPath, err, modName)
					L.RaiseError("%s", pretty)
					return 0
				}

				// Ejecutar el módulo y dejar que los errores se propaguen naturalmente
				L.Push(fn)
				L.Call(0, 1) // ←¡CAMBIAR ESTO TAMBIÉN! No usar PCall aquí

				return 1
			}

			// fallback al require original - DEJAR PROPAGAR ERROR NATURALMENTE
			L.Push(originRequire)
			L.Push(lua.LString(modName))
			L.Call(1, 1) // ← Esto está bien

			return 1
		}))
		if err := L.DoFile(scriptPath); err != nil {
			// Usar PrettyLuaError para el archivo principal también
			pretty := handler.PrettyLuaError(scriptPath, err, "") // modName vacío para el entry
			return fmt.Errorf("%s", pretty)
		}

		return nil
	},
}

func init() {
	Cmd.AddCommand(runCmd)
}
