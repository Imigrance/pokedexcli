package pokedex

type Pokedex struct {
	pokemons map[string]Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{
		pokemons: map[string]Pokemon{},
	}
}

func (p *Pokedex) AddPokemon(pokemon Pokemon) {
	if _, exists := p.pokemons[pokemon.Name]; exists {
		return
	}

	p.pokemons[pokemon.Name] = pokemon
}

func (p *Pokedex) GetPokemon(pokemon string) (Pokemon, bool) {
	if pocketmonster, exists := p.pokemons[pokemon]; exists {
		return pocketmonster, true
	}
	return Pokemon{}, false
}

func (p *Pokedex) GetAllPokemon() map[string]Pokemon {
	return p.pokemons
}

func (p *Pokedex) initMoves(pokemon *Pokemon, moves []PokeMove) {

}
