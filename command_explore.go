package main




import (
	"fmt"
)




func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid area name")
	}
	area := args[0]
	fmt.Printf("Exploring %v...\n", area)
	details, err := cfg.pokeapiClient.LocationDetails(area)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, value := range details.PokemonEncounters {
		fmt.Printf(" - %v\n", value.Pokemon.Name)
	}
	return nil
}
