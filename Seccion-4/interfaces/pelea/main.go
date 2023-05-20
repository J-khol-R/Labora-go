package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/J-khol-R/Labora-go/Seccion-4/interfaces/peleadores"
)

func main() {

	//se crean los jugadores
	var police peleadores.Police = peleadores.Police{
		BaseFighter: peleadores.BaseFighter{
			Life: 15,
		},
		Armour: 5,
	}
	var criminal peleadores.Criminal = peleadores.Criminal{
		BaseFighter: peleadores.BaseFighter{
			Life: 10,
		},
	}
	var paladin peleadores.Paladin = peleadores.Paladin{
		BaseFighter: peleadores.BaseFighter{
			Life: 200,
		},
	}

	contenders := []peleadores.Contender{&police, &criminal, &paladin}

	//se organizan aleatoriamente los contrincantes
	rand.Shuffle(len(contenders), func(i, j int) {
		contenders[i], contenders[j] = contenders[j], contenders[i]
	})

	var areAllAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
	for areAllAlive {

		random := rand.Intn(3)

		intensity := contenders[random].ThrowAttack() + 1
		fmt.Println(contenders[random].GetName(), " tira golpe con intensidad =", intensity)
		how := rand.Intn(2)
		fmt.Print(how)
		if how == 0 { //ataco a la derecha
			if random == len(contenders)-1 {
				contenders[0].RecieveAttack(intensity)
			} else if random == len(contenders)-2 {
				contenders[2].RecieveAttack(intensity)
			} else {
				contenders[1].RecieveAttack(intensity)
			}
		} else { //ataco a la izquierda
			if random == len(contenders)-1 {
				contenders[1].RecieveAttack(intensity)
			} else if random == len(contenders)-2 {
				contenders[0].RecieveAttack(intensity)
			} else {
				contenders[2].RecieveAttack(intensity)
			}
		}

		fmt.Printf("PoliceLife=%d, CriminalLife=%d, PaladinLife=%d\n", police.Life, criminal.Life, paladin.Life)
		areAllAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
		time.Sleep(3 * time.Second)
	}
}
