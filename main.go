package main




import (
	"github.com/adrake333/pokedexcli/internal/pokeapi"
	"time"
)




func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	startRepl(cfg)
}


