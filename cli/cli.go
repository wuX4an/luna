package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	debugRuntime bool
	debugBuild   bool
	versionFlag  bool
)

var Cmd = &cobra.Command{
	Use:   "luna",
	Short: "\033[1;31mLuna: A modern lua runtime",
	Long:  "\033[3;31mLuna: A modern lua runtime\033[0m",
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("luna version 0.1.0")
			return
		}
		_ = cmd.Help()
	},
}

func init() {
	Cmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Print version")

	/*
		\033[1;36m ->  Cyan Bold	Usage
		\033[1;33m ->	 Yellow Bold	Examples
		\033[1;32m ->	 Green Bold	Commands
		\033[1;34m ->	 Blue Bold	Command names
		\033[1;35m ->	 Magenta Bold	Flags
		\033[2;3m  ->  Gray italic
		\033[0m	   ->  Reset
	*/
	Cmd.SetUsageTemplate(
		// Cyan bold for "Usage"
		"\033[1;36m Usage:\033[0m {{.CommandPath}}{{if .HasAvailableSubCommands}} [command]{{end}}{{if .HasAvailableLocalFlags}} [flags]{{end}}\n\n" +

			// Yellow bold for "Examples"
			"{{- if .Example }}\n\n" +
			"\033[1;33m Examples:\033[0m\n" +
			"\033[2;3m{{.Example}}\033[0m\n\n" +
			"{{- end }}" +

			// Green bold for "Commands", and Blue bold for command names
			"{{- if .HasAvailableSubCommands }}\n\n" +
			"\033[1;32m Commands:\033[0m\n" +
			"{{range .Commands}}{{if .IsAvailableCommand}}" +
			"  \033[1;34m {{rpad .Name .NamePadding }}\033[0m {{.Short}}\n" +
			"{{end}}{{end}}" +
			"{{- end }}" +

			// Magenta bold for "Flags"
			"{{- if .HasAvailableLocalFlags }}\n" +
			"\033[1;35m Flags:\033[0m\n" +
			"{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}\n" +
			"{{- end }}",
	)

}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
