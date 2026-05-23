package main




import (
	"fmt"
	"math/rand"
)




func commandCatch(cfg *config, pokemon []string) error {
	if len(pokemon) < 1 {
		return fmt.Errorf("invalid pokemon name")
	}
	target := pokemon[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", target)
	data, err := cfg.pokeapiClient.Pokemon
}

