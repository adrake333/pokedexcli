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
	data, err := cfg.pokeapiClient.PokemonData(target)
	if err != nil {
		return err
	}
	difficulty := data.BaseExperience
	if difficulty > 450 {
		difficulty = 450
	}
	roll := rand.Intn(500)
	if roll >= difficulty {
		fmt.Printf("%v was caught!\n", target)
		cfg.pokedex[target] = data
	} else {
		fmt.Printf("%v escaped!\n", target)
	}
	return nil
}

