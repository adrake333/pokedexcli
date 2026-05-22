package main




import (
	"strings"
	"bufio"
	"fmt"
	"os"
)




func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanedLine := cleanInput(line)
		if len(cleanedLine) == 0 {
			continue
		}
		firstWord := cleanedLine[0]
		value, exists := getCommands()[firstWord]
		if exists {
			err := value.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}


type cliCommand struct {
	name 		string
	description 	string
	callback 	func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:		"exit",
			description: 	"Exit the Pokedex",
			callback: 	commandExit,
		},
		"help": {
			name:		"help",
			description: 	"Displays a help message",
			callback:	commandHelp,
		},
	}
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}



