package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Type    string
	Content string
}

func main() {

	emailChan := make(chan Message)
	smsChan := make(chan Message)
	pushChan := make(chan Message)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-emailChan:

				if !ok {
					return
				}

				fmt.Printf("Procesando mensaje de notificación correo electronico: %s\n", msg.Content)
				time.Sleep(3 * time.Second)
				fmt.Printf("se termino de procesar el mensaje correo electronico: %s\n", msg.Content)

			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-smsChan:

				if !ok {
					return
				}

				fmt.Printf("Procesando mensaje de notificación sms: %s\n", msg.Content)
				time.Sleep(3 * time.Second)
				fmt.Printf("se termino de procesar el mensaje sms: %s\n", msg.Content)

			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			select {
			case msg, ok := <-pushChan:
				if !ok {
					return
				}
				fmt.Printf("Procesando mensaje de notificación push: %s\n", msg.Content)
				time.Sleep(3 * time.Second)
				fmt.Printf("se termino de procesar el mensaje push: %s\n", msg.Content)

			}

		}
	}()

	go func(emailChan chan<- Message) {
		for i := 1; i <= 2; i++ {
			emailChan <- Message{Type: "email", Content: fmt.Sprintf("Mensaje %d de correo electrónico", i)}
		}
		close(emailChan)
	}(emailChan)

	go func(pushChan chan<- Message) {
		for i := 1; i <= 2; i++ {
			pushChan <- Message{Type: "notificación push", Content: fmt.Sprintf("Mensaje %d de notificación push", i)}
		}
		close(pushChan)
	}(pushChan)

	go func(smsChan chan<- Message) {
		for i := 1; i <= 2; i++ {
			smsChan <- Message{Type: "SMS", Content: fmt.Sprintf("Mensaje %d de SMS", i)}
		}
		close(smsChan)
	}(smsChan)

	wg.Wait()

	fmt.Println("Todos los mensajes han sido procesados.")
}
