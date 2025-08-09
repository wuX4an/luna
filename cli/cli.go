package cli

import (
	"fmt"
	"luna/cli/build"
	"luna/cli/docs"
	"luna/cli/repl"
	"luna/src/luavm"
	"os"
	"strings"

	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var (
	debugRuntime bool
	debugBuild   bool
	versionFlag  bool
)

var Cmd = &cobra.Command{
	Use:   "luna",
	Short: "\033[2;3mLuna: A modern lua runtime\033[0m",
	Long:  "\033[2;3mLuna: A modern lua runtime\033[0m",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true

		if versionFlag {
			fmt.Println("luna version 0.1.0")
			return
		}

		if len(args) == 0 {
			if err := repl.Run(); err != nil {
				fmt.Println("Error running REPL:", err)
			}
			return
		}

		arg := args[0]
		info, err := os.Stat(arg)

		// Caso 1: el argumento es un archivo existente
		if err == nil {
			if info.IsDir() {
				fmt.Fprintf(os.Stderr, "\033[31mError:\033[0m '%s' is a directory\n", arg)
				return
			}

			if strings.HasSuffix(arg, ".lua") {
				L := luavm.NewLuaVM()
				defer L.Close()

				argTable := L.NewTable()
				for i, a := range args {
					L.RawSet(argTable, lua.LNumber(i), lua.LString(a))
				}
				L.SetGlobal("arg", argTable)

				if err := L.DoFile(arg); err != nil {
					fmt.Fprintf(os.Stderr, "\033[31mError:\033[0m The module's source code could not be parsed: %v\n", err)
				}
				return
			}

			// Archivo existe pero no es .lua
			fmt.Fprintf(os.Stderr, "\033[31mError:\033[0m The module's source code could not be parsed: unsupported file type: %s\n", arg)
			return
		}

		// Caso 2: el archivo no existe
		fmt.Fprintf(os.Stderr, "\033[31mError:\033[0m Unknown argument: %s\n", arg)
		fmt.Println("\033[33mTry run\033[0m: \033[2;3mluna --help\033[0m")
	},
}

func init() {
	// Comandos
	Cmd.AddCommand(build.BuildCmd)
	Cmd.AddCommand(docs.DocCmd)
	Cmd.AddCommand(repl.ReplCmd)

	// Flags
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
			"{{- if .HasAvailableLocalFlags }}" +
			"\n\033[1;35m Flags:\033[0m\n" +
			"{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}\n" +
			"{{- end }}\n",
	)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
