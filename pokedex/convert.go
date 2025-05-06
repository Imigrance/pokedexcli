package pokedex

func ConvertRespToPokemon(response RespPokemon) Pokemon {
	pokemon := Pokemon{}

	pokemon.Abilities = map[string]Ability{}
	pokemon.BaseExperience = response.BaseExperience
	pokemon.Height = response.Height
	pokemon.HeldItems = map[string]HeldItem{}
	pokemon.ID = response.ID
	pokemon.IsDefault = response.IsDefault
	pokemon.Level = 5
	pokemon.LearnedMoves = map[string]PokeMove{}
	pokemon.LocationAreaEncounters = response.LocationAreaEncounters
	pokemon.Moves = map[string]Move{}
	pokemon.Name = response.Name
	pokemon.Order = response.Order
	pokemon.PastAbilities = response.PastAbilities
	pokemon.PastTypes = response.PastTypes
	pokemon.Species = struct {
		Name string
		URL  string
	}(response.Species)
	pokemon.Stats = map[string]PokeStats{}
	pokemon.Types = map[string]PokeType{}
	pokemon.Weight = response.Weight

	for _, a := range response.Abilities {
		pokemon.Abilities[a.Ability.Name] = Ability{
			Name:     a.Ability.Name,
			URL:      a.Ability.URL,
			IsHidden: a.IsHidden,
			Slot:     a.Slot,
		}
	}

	for _, i := range response.HeldItems {
		pokemon.HeldItems[i.Item.Name] = HeldItem{
			Name: i.Item.Name,
			URL:  i.Item.URL,
		}
	}

	for _, m := range response.Moves {
		move := Move{}
		move.Name = m.Move.Name
		move.URL = m.Move.URL

		for _, d := range m.VersionGroupDetails {
			move.LevelLearnedAt = d.LevelLearnedAt
			move.MoveLearnMethod = struct {
				Name string
				URL  string
			}(d.MoveLearnMethod)

			break
		}

		pokemon.Moves[m.Move.Name] = move
	}

	for _, s := range response.Stats {
		pokemon.Stats[s.Stat.Name] = PokeStats{
			BaseStat: s.BaseStat,
			Effort:   s.Effort,
			Name:     s.Stat.Name,
			URL:      s.Stat.URL,
		}
	}

	pokemon.Stats["maxhp"] = PokeStats{
		BaseStat: pokemon.Stats["hp"].BaseStat,
		Effort:   pokemon.Stats["hp"].Effort,
		Name:     "Max HP",
		URL:      pokemon.Stats["hp"].URL,
	}

	for _, t := range response.Types {
		pokemon.Types[t.Type.Name] = PokeType{
			Slot: t.Slot,
			Type: t.Type.Name,
			URL:  t.Type.URL,
		}
	}

	return pokemon
}
