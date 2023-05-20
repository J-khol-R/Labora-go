package peleadores

import "math/rand"

type Police struct {
	BaseFighter
	Armour int // 0..50
}

func (p Police) ThrowAttack() int {
	return rand.Intn(4)
}

func (p *Police) RecieveAttack(intensity int) {
	if p.Armour > 0 {
		diff := (p.Armour - intensity)
		hasArmour := diff > 0
		if hasArmour {
			p.Armour -= intensity
			intensity = 0
		} else {
			p.Armour = 0
			intensity = -diff // intensity -= p.Armour
		}
	}
	p.Life -= intensity
}

func (p Police) GetName() string {
	return "Policia"
}
