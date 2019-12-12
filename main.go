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

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input := ""
	output := []string{}
	cmdRun := false

	trieObj := trie.NewTrie()

	// infinite loop
	for {
		if !cmdRun {
			fmt.Printf("> %s", input)
			fmt.Printf("> Output: %s", output[0:])

		}

		// errors coming from user's shell
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		// TODO: check if input already exists in $PATH
		// if user hits 'enter' key with chars
		if char == '\n' && len(input) > 0 {
			// running commands and checking for errors
			if err = execInput(input); err != nil {
				fmt.Fprint(os.Stderr, err)
				input = ""
				output = []string{""}
				continue
			}

			cmdRun = true
			// since no error, insert the input
			trieObj.Insert(input)
			input = ""
		} else {
			input = input + string(char)
			output = trieObj.Complete(input)

			cmdRun = false
		}

	}

}
