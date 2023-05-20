package peleadores

import "math/rand"

type Paladin struct {
	BaseFighter
}

func (p *Paladin) ThrowAttack() int {
	ataque := rand.Intn(10) + 1
	proporcion := p.Life / 200
	return ataque * proporcion
}

func (p *Paladin) RecieveAttack(intensity int) {
	p.Life -= intensity
}

func (p *Paladin) GetName() string {
	return "Paladin"
}
