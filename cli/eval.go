package cli

import (
	"fmt"
	"luna/src/luavm"

	"github.com/spf13/cobra"
)

var evalExamples = `
  luna eval 'print("Hello, Luna!")'
  luna eval 'for i=1,3 do print(i) end'`

var evalCmd = &cobra.Command{
	Use:     "eval [CODE]",
	Short:   "Evaluate a Lua snippet",
	Example: evalExamples,
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		code := args[0]
		L := luavm.NewLuaVM()
		defer L.Close()

		if err := L.DoString(code); err != nil {
			return fmt.Errorf("\n  Error evaluating code: %w", err)
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(evalCmd)
}
