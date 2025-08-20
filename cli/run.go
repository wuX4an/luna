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

		// Run main script
		if err := L.DoFile(scriptPath); err != nil {
			return fmt.Errorf("error running script: %w", err)
		}

		return nil
	},
}

func init() {
	Cmd.AddCommand(runCmd)
}
