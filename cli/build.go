package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	buildTarget string
	buildArch   string
)

var buildExamples = `
  luna build main.ts
  luna build .
  luna build . --target=linux --arch=arm64
`

var buildCmd = &cobra.Command{
	Use:     "build [FILE]",
	Short:   "Compile the script into a self contained executable",
	Example: buildExamples, Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := ""
		if len(args) > 0 {
			file = args[0]
		}
		fmt.Printf("Building %q for target=%q arch=%q\n", file, buildTarget, buildArch)
		return nil
	},
}

func init() {
	Cmd.AddCommand(buildCmd)
	buildCmd.Flags().StringVar(&buildTarget, "target", "", "Target OS (linux/darwin/windows)")
	buildCmd.Flags().StringVar(&buildArch, "arch", "", "Target architecture (amd64/arm64)")
}
