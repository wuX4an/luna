package docs

import (
	_ "embed"
	"fmt"
	"luna/cli/docs/ui"

	"github.com/spf13/cobra"
)

//go:embed manual/index.txt
var IndexContent string

var (
	docDir string
)

var DocCmd = &cobra.Command{
	Use:     "docs",
	Short:   "View Luna's documentation",
	Example: "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("the 'docs' command does not accept any arguments")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Viewing docs")
		ui.Tui()
		return nil
	},
}
