package peleadores

import "math/rand"

type Criminal struct {
	BaseFighter
}

func (p *Criminal) ThrowAttack() int {
	return rand.Intn(8)
}

func (p *Criminal) RecieveAttack(intensity int) {
	p.Life -= intensity
}

func (p *Criminal) GetName() string {
	return "Criminal"
}
