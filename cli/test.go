package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test [FILE]",
	Short:   "Run tests",
	Example: "luna test main.lua\nluna test .",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := ""
		if len(args) > 0 {
			file = args[0]
		}
		fmt.Printf("Running tests for %q\n", file)
		return nil
	},
}

func init() {
	Cmd.AddCommand(testCmd)
}
