package cli

import (
	"fmt"
	"io/fs"
	"luna/src/luavm"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var runExamples = `  luna run main.lua
  luna run .`

var runCmd = &cobra.Command{
	Use:     "run [FILE|DIR]",
	Short:   "Run a lua program or project",
	Example: runExamples,
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		var scriptPath string
		var projectDir string

		// Si no hay argumento, asumimos directorio actual
		if len(args) == 0 {
			args = append(args, ".")
		}
		arg := args[0]

		info, err := os.Stat(arg)
		if err != nil {
			return fmt.Errorf("path does not exist: %s", arg)
		}

		if info.IsDir() {
			projectDir = arg

			// Buscar Luna.toml para entry
			configPath := filepath.Join(arg, "Luna.toml")
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				return fmt.Errorf("could not find Luna.toml in %s", arg)
			}

			configData, err := os.ReadFile(configPath)
			if err != nil {
				return fmt.Errorf("could not read Luna.toml: %w", err)
			}

			var conf struct {
				Build struct {
					Entry  string `toml:"entry"`  // archivo principal
					Source string `toml:"source"` // directorio de módulos
				} `toml:"build"`
			}
			if err := toml.Unmarshal(configData, &conf); err != nil {
				return fmt.Errorf("invalid Luna.toml: %w", err)
			}

			if conf.Build.Entry == "" {
				return fmt.Errorf("Luna.toml missing [build].entry")
			}
			if conf.Build.Source == "" {
				return fmt.Errorf("Luna.toml missing [build].source")
			}

			scriptPath = filepath.Join(arg, conf.Build.Source, conf.Build.Entry)
			projectDir = filepath.Join(arg, conf.Build.Source)

		} else if filepath.Ext(arg) == ".lua" {
			// Archivo individual
			scriptPath = arg
			projectDir = filepath.Dir(arg)
		} else {
			return fmt.Errorf("argument must be a .lua file or a project directory with Luna.toml")
		}

		// Inicializar VM Lua
		L := luavm.NewLuaVM()
		defer L.Close()

		// Cargar todos los módulos en memoria
		modules := map[string]string{}
		err = filepath.WalkDir(projectDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			if filepath.Ext(path) == ".lua" {
				rel, _ := filepath.Rel(projectDir, path)
				rel = filepath.ToSlash(rel) // unix-style: "hello.lua" o "foo/bar.lua"
				base := strings.TrimSuffix(rel, ".lua")

				// Nombre con ":" (tu estilo)
				modName := strings.ReplaceAll(base, "/", ":")

				data, _ := os.ReadFile(path)

				// Guardar tanto con ":" como con "." como alias
				modules[modName] = string(data)
				modules[strings.ReplaceAll(base, "/", ".")] = string(data)
			}
			return nil
		})

		// Redefinir require para buscar en memoria
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

		// Ejecutar script principal
		if err := L.DoFile(scriptPath); err != nil {
			return fmt.Errorf("error running script: %w", err)
		}

		return nil
	},
}

func init() {
	Cmd.AddCommand(runCmd)
}
