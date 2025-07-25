package cli

import (
	"fmt"

	"github.com/spf13/cobra"
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
		file := ""
		if len(args) > 0 {
			file = args[0]
		}
		fmt.Printf("Running %q\n", file)
		return nil
	},
}

func init() {
	Cmd.AddCommand(runCmd)
}
