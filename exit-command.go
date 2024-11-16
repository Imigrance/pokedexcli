package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {

	err := saveGame(cfg)
	if err != nil {
		fmt.Printf("Error saving game: %v", err)
	}

	os.Exit(0)
	return nil
}
