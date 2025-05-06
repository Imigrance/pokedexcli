package pokedex

func (p *Pokemon) TakeDmg(dmg int) {
	if stat, ok := p.Stats["hp"]; ok {
		stat.BaseStat -= dmg
		p.Stats["hp"] = stat
	}
}
