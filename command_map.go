package main

import (
	"fmt"
)




func commandMap(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locations.Next
	cfg.previousLocationsURL = locations.Previous
	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}


func commandMapb(cfg *config) error {
	if cfg.previousLocationsURL == nil {
		fmt.Printlm("you're on the first page")
		return nil
	}
	locations, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locations.Next
	cfg.previousLocationsURL = locations.Previous
	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

