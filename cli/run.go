package cli

import (
	"fmt"
	"luna/std"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var runExamples = `
  luna run main.lua
  luna run .
`

var runCmd = &cobra.Command{
	Use:     "run [FILE]",
	Short:   "Run a lua program, or project",
	Example: runExamples,
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true  // <- siempre silenciar ayuda automática
		cmd.SilenceErrors = true // <- siempre silenciar mensaje "Error: ..."

		if len(args) == 0 {
			_ = cmd.Help()
			return fmt.Errorf("\nError: No script specified")
		}

		L := lua.NewState()
		defer L.Close()

		std.RegisterAll(L)

		script := args[0]
		if err := L.DoFile(script); err != nil {
			return fmt.Errorf("\n  Error running script: %w", err)
		}

		return nil
	},
}

func init() {
	Cmd.AddCommand(runCmd)
}
