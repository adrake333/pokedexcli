package main




import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

import "github.com/adrake333/pokedexcli/internal/pokeapi"




func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}


func startRepl() {
	cfg := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		cleanedLine := cleanInput(line)
		if len(cleanedLine) == 0 {
			continue
		}
		firstWord := cleanedLine[0]
		value, exists := commands[firstWord]
		if exists {
			err := value.callback(cfg)
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
	callback 	func(*config) error
}


type config struct {
	nextLocationsURL	*string
	previousLocationsURL	*string
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
		"map": {
			name:		"map",
			description:	"Displays the next 20 location areas in the Pokemon World",
			callback:	commandMap,
		},
		"mapb": {
			name:		"mapb",
			description:	"Displays the previous 20 location areas in the Pokemon World",
			callback:	commandMapb,
		},
	}
}






