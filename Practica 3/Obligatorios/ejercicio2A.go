package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func cargarCliente(in chan string, wg *sync.WaitGroup) {

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i := 1; i < 10; i++ {

		b := make([]rune, 7)

		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		in <- string(b)
	}
	wg.Done()
}

func simularAtencion(caja chan string, clientes chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		cli, ok := <-clientes
		if !ok {
			break // El canal se cerró, salir del bucle
		}

		// Crear un temporizador para cada cliente
		tiempoAtencion := time.Duration(rand.Intn(1000)) * time.Millisecond

		caja <- cli
		fmt.Println("Cliente:", cli, " siendo atendido")
		time.Sleep(tiempoAtencion)
		fmt.Println("El cliente:", <-caja, " fue atendido")
	}
	close(caja)

}

func main() {
	// Channels
	colaClientes := make(chan string)
	caja := make(chan string, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	// Agregar clientes
	go cargarCliente(colaClientes, &wg)

	// Simular atención
	go simularAtencion(caja, colaClientes, &wg)

	wg.Wait()

	// Esperar a que simularAtencion termine de procesar todos los clientes
	close(colaClientes)

	fmt.Println("Finaliza la atención")
}
