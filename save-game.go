package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/imigrance/pokedexcli/pokedex"
)

const (
	SAVE_FOLDER   = "save"
	SAVE_FILENAME = "save.json"
)

func saveGame(cfg *config) error {
	ok := saveExists()
	if !ok {
		isFolder := isSaveFolderAvail()
		if !isFolder {
			fmt.Println("Creating save folder!")
			os.Mkdir(SAVE_FOLDER, os.ModePerm)
		}
	}

	data, err := json.Marshal(cfg.pokedex.GetAllPokemon())
	if err != nil {
		return err
	}

	os.WriteFile(SAVE_FOLDER+"/"+SAVE_FILENAME, data, os.ModePerm)
	return nil
}

func loadSave(cfg *config) error {
	ok := saveExists()
	if !ok {
		return errors.New("no save available")
	}

	file, err := os.ReadFile(SAVE_FOLDER + "/" + SAVE_FILENAME)
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

	fmt.Println(">>> Save loaded <<<")
	return nil
}

func saveExists() bool {
	ok := isSaveFolderAvail()
	if !ok {
		return false
	}

	_, err := os.ReadFile(SAVE_FOLDER + "/" + SAVE_FILENAME)
	return err == nil
}

func isSaveFolderAvail() bool {
	_, err := os.ReadDir(SAVE_FOLDER)
	return err == nil
}
