package battle

type battleCommand struct {
	name        string
	description string
	callback    func(*battleConfig, ...string) error
}

func getBattleCommands() map[string]battleCommand {
	return map[string]battleCommand{
		"attack": {
			name:        "attack",
			description: "Attack with move",
			callback:    battleCommandAttack,
		},
		"use-item": {
			name:        "use-item",
			description: "Use an item in battle",
			callback:    battleCommandUseItem,
		},
		"flee": {
			name:        "flee",
			description: "Tries to flee from current battle",
			callback:    batteCommandFlee,
		},
		"switch": {
			name:        "switch",
			description: "Switches to another Pokemon",
			callback:    battleCommandSwitch,
		},
	}
}
