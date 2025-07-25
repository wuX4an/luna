package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [DIR]",
	Short:   "Initialize a new project",
	Example: "luna init\nluna init hello-world",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}
		fmt.Printf("Initializing project in %q\n", dir)
		return nil
	},
}

func init() {
	Cmd.AddCommand(initCmd)
}
