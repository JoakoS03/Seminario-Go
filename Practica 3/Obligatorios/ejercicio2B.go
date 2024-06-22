package main

import (
	"fmt"
	"math/rand"
	"time"
)

func caja(id int, clientes <-chan string, terminado chan<- int) {
	for cliente := range clientes {
		tiempoAtencion := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("Caja %d atendiendo cliente %s\n", id, cliente)
		time.Sleep(tiempoAtencion)
		fmt.Printf("Caja %d terminó de atender cliente %s\n", id, cliente)
		terminado <- 1
	}
}

func main() {
	numCajas := 3
	numClientes := 10

	clientes := make(chan string, numClientes)
	terminado := make(chan int)

	for i := 0; i < numClientes; i++ {
		clientes <- fmt.Sprintf("Cliente%d", i+1)
	}

	for i := 0; i < numCajas; i++ {
		go caja(i+1, clientes, terminado)
	}

	close(clientes)

	atendidos := 0
	for atendidos < numClientes {
		select {
		case <-terminado:
			atendidos++
		}
	}
	fmt.Println(atendidos)
}
