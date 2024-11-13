package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("missing area")
	}

	inArea, err := cfg.pokeapiClient.GetInArea(args[0])
	if err != nil {
		return err
	}

	encounters := inArea.PokemonEncounters

	fmt.Printf("Pokemon found in %v\n", args[0])
	for _, encounter := range encounters {
		fmt.Printf("- %v \n", encounter.Pokemon.Name)
	}

	return nil
}
