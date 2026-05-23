 package main




 import (
	 "fmt"
 )




func commandPokedex(cfg *config, _ []string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Printf(" - %v\n", name)
	}
	return nil
}
