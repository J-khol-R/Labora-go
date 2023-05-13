package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var balance = 1000
var mutex sync.Mutex

func transaccion(tipo int, wg *sync.WaitGroup, m *sync.Mutex) {
	rand.Seed(time.Now().UnixNano())
	monto := rand.Intn(500) + 1
	m.Lock()
	defer m.Unlock()
	// si es par es de tipo retiro y si es impar de tipo depositos
	if (tipo % 2) == 0 {
		if balance < monto {
			fmt.Printf("fondos insuficientes total en la cuenta: %d", balance)
			return
		} else {
			fmt.Println("restar")
			fmt.Println(monto)
			balance -= monto
		}
	}
	if (tipo % 2) == 1 {
		fmt.Println("sumar")
		fmt.Println(monto)
		balance += monto
	}
	fmt.Println(balance)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 9; i++ {
		wg.Add(1)
		go transaccion(i, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println(balance)

}
