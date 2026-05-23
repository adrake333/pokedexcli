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
		for _, stat := range value.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, poketype := range value.Types {
			fmt.Printf("  - %v\n", poketype.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
