package main




import (
	"fmt"
)




func commandInspect(cfg *config, pokemon []string) error {
	if len(pokemon) < 1 {
		return fmt.Errorf("invalid pokemon name")
	}
	target := pokemon[0]
	if value, ok := cfg.pokedex[target]; ok {
		fmt.Printf("Name: %v\n", value.Name)
		fmt.Printf("Height: %v\n", value.Height)
		fmt.Printf("Weight: %v\n", value.Weight)
		fmt.Println("Stats:")
		fmt.Printf("  -hp: %v\n", value.//???)
