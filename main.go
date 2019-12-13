// Thanks to @sj14 for providing a tutorial on how to create a simple shell in Go!
package main

import (
	"bufio"
	"fmt"
	"github.com/ThomasLee94/autosuggest/trie"
	"github.com/gookit/color"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	// array to take in multiple commands
	args := strings.Split(input, " ")

	switch args[0] {
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input := ""
	// output that will be filled with autosuggests after insertion
	output := []string{}
	cmdRun := false

	trieObj := trie.NewTrie()

	// infinite loop
	for {
		fmt.Print("$ ")
		// case: no command yet given.
		if !cmdRun {
			fmt.Printf("> %s\n", input)
			fmt.Printf("> Output: %s\n", output[0:])
			// colour output
			color.Cyan.Printf("Ouput: %s\n", output[0:])
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
			err = execInput(input)
			if err != nil {
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
