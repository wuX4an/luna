package cli

import (
	"fmt"
	"luna/src/config"
	"luna/src/luavm"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task [NAME]",
	Short: "Run or list tasks",
	Example: `  luna task             # list all tasks
  luna task build       # run the "build" task
  luna task run         # run the "run" task`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := config.LoadConfig("Luna.toml")
		if err != nil {
			return err
		}

		// si no pasaron nombre, listar
		if len(args) == 0 {
			fmt.Println("Available tasks:")
			for name, td := range conf.Tasks.Defs {
				desc := td.Desc
				if desc == "" {
					desc = td.Script
				}
				fmt.Printf("  %s\t%s\n", name, desc)
			}
			return nil
		}

		taskName := args[0]
		visited := map[string]bool{}
		return runTask(taskName, conf, visited)
	},
}

func runTask(name string, conf *config.LunaConfig, visited map[string]bool) error {
	if visited[name] {
		return nil // ya corrido
	}
	td, ok := conf.Tasks.Defs[name]
	if !ok {
		return fmt.Errorf("task %q not found", name)
	}

	// primero correr las dependencias
	for _, dep := range td.Depends {
		if err := runTask(dep, conf, visited); err != nil {
			return fmt.Errorf("while running dependency %q of %q: %w", dep, name, err)
		}
	}

	// ahora correr el task actual
	scriptPath := filepath.Join(conf.Tasks.Source, td.Script)
	if _, err := os.Stat(scriptPath); err != nil {
		return fmt.Errorf("cannot find task script: %s", scriptPath)
	}

	fmt.Printf("▶ running task %s (%s)\n", name, td.Script)

	L := luavm.NewLuaVM()
	defer L.Close()

	if err := L.DoFile(scriptPath); err != nil {
		return fmt.Errorf("error running %s: %w", scriptPath, err)
	}

	visited[name] = true
	return nil
}

func init() {
	Cmd.AddCommand(taskCmd)
}
