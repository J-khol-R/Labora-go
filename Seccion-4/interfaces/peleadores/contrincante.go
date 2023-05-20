package peleadores

type Contender interface {
	ThrowAttack() int
	RecieveAttack(intensity int)
	IsAlive() bool
	GetName() string
}
