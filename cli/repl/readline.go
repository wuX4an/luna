package repl

import (
	"regexp"
	"strings" // Import the strings package

	"github.com/chzyer/readline"
)

// Do se llama cuando el usuario pulsa TAB, para sugerencias
var wordRegexp = regexp.MustCompile(`[\w_]+$`)

func (c *LuaAutoCompleter) Do(line []rune, pos int) ([][]rune, int) {
	input := string(line[:pos])
	word := wordRegexp.FindString(input) // 'word' is the current incomplete word

	// Completions are generated based on 'word'
	completions := c.repl.Completer(word)

	suggestions := make([][]rune, 0, len(completions))
	for _, comp := range completions {
		// Calculate the suffix to append
		// If "print" is the completion and "pri" is the word, then the suffix is "nt"
		if strings.HasPrefix(comp, word) {
			suffix := comp[len(word):]
			suggestions = append(suggestions, []rune(suffix))
		} else {
			// This case should ideally not happen if Completer filters correctly
			// but it's a safeguard.
			suggestions = append(suggestions, []rune(comp))
		}
	}

	// The crucial part: we still tell readline how many characters
	// from the end of the input string should be replaced.
	// This should be the length of the 'word' we found, as we're now providing the suffix.
	return suggestions, len([]rune(word))
}

func NewReadline(repl *LuaRepl) (*readline.Instance, error) {
	completer := &LuaAutoCompleter{repl: repl}
	rl, err := readline.NewEx(&readline.Config{
		// Prompt in Bright Green
		Prompt:            "\033[92m> \033[0m",
		HistoryFile:       "/tmp/luna_repl_history.tmp",
		InterruptPrompt:   "^D",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
		AutoComplete:      completer, // <- aquí debes asignarlo directamente
	})
	if err != nil {
		return nil, err
	}
	return rl, nil
}
