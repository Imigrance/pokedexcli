package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	for _, res := range locations.Results {
		fmt.Println(res.Name)
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return fmt.Errorf("there is only one way -> forward")
	}

	locations, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	for _, res := range locations.Results {
		fmt.Println(res.Name)
	}

	cfg.nextLocationURL = locations.Next
	cfg.prevLocationURL = locations.Previous

	return nil
}
