package main

type CliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]CliCommand {
	commands := map[string]CliCommand{
		"battle": {
			name:        "battle",
			description: "	Battles 2 pokemons",
			callback:    commandBattle,
		},
		"catch": {
			name:        "catch",
			description: "		Attempts to catch a specified pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "		Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "	Shows pokemon found in the area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "		Displays help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "	Inspects a pokemon in the pokedex",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "		Gets the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "		Gets the previous 20 locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "	Shows all the pokemon in your pokedex",
			callback:    commandPokedex,
		},
	}
	return commands
}
