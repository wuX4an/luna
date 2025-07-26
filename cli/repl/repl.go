package repl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Run() error {
	// Enhanced Welcome Banner Colors
	fmt.Println("\033[1;34m       -------------------------------------------\033[0m")                    // Bold Blue border
	fmt.Println("\033[1;34m      | \033[97m           July 20, 1969 A.D.             \033[1;34m|\033[0m") // Bright White text
	fmt.Println("\033[1;34m      | \033[97m     Here Men from the Planet Earth       \033[1;34m|\033[0m")
	fmt.Println("\033[1;34m      | \033[97m      First set Foot upon the Moon        \033[1;34m|\033[0m")
	fmt.Println("\033[1;34m      | \033[97m    We came in Peace for all Mankind      \033[1;34m|\033[0m")
	fmt.Println("\033[1;34m       ------------------------------== \033[1;95mluna\033[0m \033[1;34m==---\033[0m") // Bold Magenta for "luna"
	fmt.Println("")                                                                                               // Add a newline for spacing

	repl := NewLuaRepl()
	defer repl.Close()

	rl, err := NewReadline(repl)
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // EOF o Ctrl+C
			// Exit message in Bold Yellow
			fmt.Println("\n\033[1;33mMission complete. Safe travels, astronaut.\033[0m")
			break
		}
		output := repl.Eval(line)
		if output != "" {
			fmt.Print(output)
		}
	}
	return nil
}

var ReplCmd = &cobra.Command{
	Use:   "repl",
	Short: "Start an interactive Lua Read-Eval-Print Loop (REPL)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Run()
	},
}
