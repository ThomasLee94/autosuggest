package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ThomasLee94/autosuggest/trie"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	// path := os.Getenv("PATH")

	// cmd := exec.Cmd{
	// 	Path:   path,
	// 	Args:   args,
	// 	Env:    []string{fmt.Sprintf("PATH=%s", path)},
	// 	Stdout: os.Stdout,
	// 	Stderr: os.Stderr,
	// }

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input := ""
	output := ""
	cmdRun := false

	trie := trie.NewTrie()

	for {
		if !cmdRun {
			fmt.Printf("> %s", input)
			fmt.Printf("> Output: %s", output)
		}

		// errors coming from user's shell
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		// TODO: check if input already exists in $PATH
		if char == '\n' && len(input) > 0 {
			// running commands and checking for errors
			if err = execInput(input); err != nil {
				fmt.Fprint(os.Stderr, err)
				input = ""
				output = ""
				continue
			}

			cmdRun = true
			// since no error, insert the input
			trie.Insert(input)
			input = ""
		} else {
			input = input + string(char)
			output = trie.Complete(input)

			cmdRun = false
		}

	}

}
