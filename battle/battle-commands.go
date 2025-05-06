package battle

type battleCommand struct {
	name        string
	description string
	callback    func(*battleConfig, ...string) error
}

func getBattleCommands() map[string]battleCommand {
	return map[string]battleCommand{
		"use-move": {
			name:        "use-move",
			description: "Use a move to attack the other pokemon or buff your own",
			callback:    battleCommandUseMove,
		},
		/*"use-item": {
			name:        "use-item",
			description: "Use an item in battle",
			callback:    battleCommandUseItem,
		},
		*/"flee": {
			name:        "flee",
			description: "Tries to flee from current battle",
			callback:    battleCommandFlee,
		},
		/*"switch": {
			name:        "switch",
			description: "Switches to another Pokemon",
			callback:    battleCommandSwitch,
		},*/
	}
}
