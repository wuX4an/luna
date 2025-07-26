package docs

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
		return nil
	},
}
