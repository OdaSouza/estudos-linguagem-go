package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("🏓 PING-PONG CONTROLADO 🏓")
	fmt.Println("==========================")

	var wg sync.WaitGroup
	pingChan := make(chan string)
	pongChan := make(chan string)
	done := make(chan bool)

	// Contador de jogadas
	jogadas := 0
	maxJogadas := 10

	// Goroutine do Ping
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case msg := <-pingChan:
				jogadas++
				fmt.Printf("Ping recebeu: %s (Jogada %d)\n", msg, jogadas)

				if jogadas >= maxJogadas {
					fmt.Println("🏆 Ping venceu! Jogo finalizado.")
					close(done)
					return
				}

				time.Sleep(300 * time.Millisecond)
				pongChan <- "🏓 PING!"

			case <-done:
				return
			}
		}
	}()

	// Goroutine do Pong
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case msg := <-pongChan:
				fmt.Printf("Pong recebeu: %s (Jogada %d)\n", msg, jogadas)

				time.Sleep(300 * time.Millisecond)
				pingChan <- "🏸 PONG!"

			case <-done:
				return
			}
		}
	}()

	// Inicia o jogo
	fmt.Printf("Jogo começando! Máximo de %d jogadas.\n", maxJogadas)
	pingChan <- "🎾 SAQUE!"

	// Aguarda término
	wg.Wait()
	fmt.Println("🎯 Jogo finalizado!")
}
