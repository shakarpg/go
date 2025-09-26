package main

import (
	"fmt"
	"time"
)

func ping(pingChan chan<- string, pongChan <-chan string) {
	for {
		// Envia "ping" para o canal
		pingChan <- "ping"

		// Aguarda receber sinal do pong para continuar
		<-pongChan

		// Pequena pausa para visualizar melhor a alternância
		time.Sleep(500 * time.Millisecond)
	}
}

func pong(pongChan chan<- string, pingChan <-chan string) {
	for {
		// Aguarda receber "ping" do canal
		message := <-pingChan
		fmt.Println(message)

		// Envia "pong" para o canal
		pongChan <- "pong"

		// Imprime pong
		fmt.Println("pong")

		// Pequena pausa para visualizar melhor a alternância
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Criando canais para comunicação entre goroutines
	pingChan := make(chan string)
	pongChan := make(chan string)

	fmt.Println("Iniciando o jogo Ping Pong com goroutines!")
	fmt.Println("Pressione Ctrl+C para parar...")

	// Iniciando as goroutines
	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)

	// Mantém o programa principal rodando
	// Usando select vazio para bloquear indefinidamente
	select {}
}
