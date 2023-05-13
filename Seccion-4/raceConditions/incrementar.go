package main

import (
	"fmt"
	"sync"
)

var valor int
var mutex sync.Mutex

func incrementar(wg *sync.WaitGroup, m *sync.Mutex) {
	for i := 0; i < 100; i++ {
		m.Lock()
		valor++
		m.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go incrementar(&wg, &mutex)
	go incrementar(&wg, &mutex)
	wg.Wait()

	fmt.Println(valor)
}
