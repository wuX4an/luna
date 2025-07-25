package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Start an interactive Read-Eval-Print Loop (REPL)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting interactive REPL...")
		return nil
	},
}

func init() {
	Cmd.AddCommand(replCmd)
}
