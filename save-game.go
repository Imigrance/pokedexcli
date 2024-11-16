package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/imigrance/pokedexcli/pokedex"
)

func saveGame(cfg *config) error {
	//os.WriteFile("save.json", cfg.pokedex, 0666)

	ok := saveExists()
	if !ok {
		isFolder := isSaveFolderAvail()
		if !isFolder {
			fmt.Println("Creating save folder!")
			os.Mkdir("save", os.ModePerm)
		}
	}

	data, err := json.Marshal(cfg.pokedex.GetAllPokemon())
	if err != nil {
		return err
	}

	os.WriteFile("./save/save.json", data, os.ModePerm)
	return nil
}

func loadSave(cfg *config) error {
	ok := saveExists()
	if !ok {
		return errors.New("no save available")
	}

	file, err := os.ReadFile("./save/save.json")
	if err != nil {
		return err
	}

	data := map[string]pokedex.Pokemon{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	for _, pokemon := range data {
		cfg.pokedex.AddPokemon(pokemon)
	}

	fmt.Println("Save loaded")
	return nil
}

func saveExists() bool {
	ok := isSaveFolderAvail()
	if !ok {
		return false
	}

	_, err := os.ReadFile("./save/save.json")
	return err == nil
}

func isSaveFolderAvail() bool {
	_, err := os.ReadDir("save")
	return err == nil
}
