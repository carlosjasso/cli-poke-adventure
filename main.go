package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/carlosjasso/cli-poke-adventure/prompt"
)

func main() {
	for {
		prompt.Display()

		stdIn := bufio.NewScanner(os.Stdin)
		stdIn.Scan()
		input := stdIn.Text()

		terms := strings.Split(input, " ")
		commandName := prompt.CommandName(terms[0])
		command := prompt.Commands[commandName]

		if command == nil {
			fmt.Println("Invalid command")
		} else if command.Name == prompt.COMMAND_EXIT {
			break
		} else {
			command.Action()
		}
	}
}
