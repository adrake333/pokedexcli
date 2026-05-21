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
		fmt.Printf("Your command was: %v\n", firstWord)
	}
}
