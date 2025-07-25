package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "clean [DIR]",
	Short:   "Remove the *build* directory",
	Example: "luna clean\nluna clean .",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "build"
		if len(args) > 0 {
			dir = args[0]
		}
		fmt.Printf("Cleaning directory %q\n", dir)
		return nil
	},
}

func init() {
	Cmd.AddCommand(cleanCmd)
}
