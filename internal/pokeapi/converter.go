package pokeapi

import "github.com/imigrance/pokedexcli/pokedex"

func ConvertRespToPokeMove(resp RespPokeMove) pokedex.PokeMove {
	pokeMove := pokedex.PokeMove{}

	pokeMove.Accuracy = resp.Accuracy
	pokeMove.DamageClass = struct {
		Name string
		URL  string
	}(resp.DamageClass)
	pokeMove.EffectChance = resp.EffectChance
	pokeMove.EffectChanges = []interface{}(resp.EffectChanges)
	pokeMove.EffectEntries = map[string]pokedex.MoveEffect{}
	pokeMove.ID = resp.ID
	pokeMove.Meta = struct {
		Ailment struct {
			Name string
			URL  string
		}
		AilmentChance int
		Category      struct {
			Name string
			URL  string
		}
		CritRate     int
		Drain        int
		FlinchChance int
		Healing      int
		MaxHits      interface{}
		MaxTurns     interface{}
		MinHits      interface{}
		MinTurns     interface{}
		StatChance   int
	}(resp.Meta)
	pokeMove.Name = resp.Name
	pokeMove.PastValues = resp.PastValues
	pokeMove.Power = resp.Power
	pokeMove.Pp = resp.Pp
	pokeMove.Priority = resp.Priority
	pokeMove.StatChanges = resp.StatChanges
	pokeMove.Target = struct {
		Name string
		URL  string
	}(resp.Target)
	pokeMove.Type = struct {
		Name string
		URL  string
	}(resp.Type)

	for _, effect := range resp.EffectEntries {
		pokeMove.EffectEntries[effect.Effect] = pokedex.MoveEffect{
			Effect:      effect.Effect,
			ShortEffect: effect.ShortEffect,
		}
	}

	for _, txt := range resp.FlavorTextEntries {
		if txt.Language.Name == "en" {
			pokeMove.FlavorText = txt.FlavorText
			break
		}
	}

	return pokeMove
}
