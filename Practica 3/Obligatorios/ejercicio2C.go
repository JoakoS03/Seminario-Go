package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Caja struct {
	id       int
	clientes chan string
}

func nuevaCaja(id int) Caja {
	return Caja{
		id:       id,
		clientes: make(chan string, 10), // Capacidad de la cola de cada caja
	}
}

func (c *Caja) atenderClientes(wg *sync.WaitGroup) {
	defer wg.Done()
	for cliente := range c.clientes {
		tiempoAtencion := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("Caja %d atendiendo cliente %s\n", c.id, cliente)
		time.Sleep(tiempoAtencion)
		fmt.Printf("Caja %d terminó de atender cliente %s\n", c.id, cliente)
	}
}

func cajaConColaMasCorta(cajas []Caja) *Caja {
	var cajaConMenosClientes *Caja
	minClientes := int(^uint(0) >> 1) // Inicializa con el valor máximo de int
	for i := range cajas {
		if len(cajas[i].clientes) < minClientes {
			cajaConMenosClientes = &cajas[i]
			minClientes = len(cajas[i].clientes)
		}
	}
	return cajaConMenosClientes
}

func main() {

	numCajas := 3
	numClientes := 10

	var wg sync.WaitGroup

	// Crear un slice de estructuras Caja
	cajas := make([]Caja, numCajas)

	for i := 0; i < numCajas; i++ {
		cajas[i] = nuevaCaja(i + 1)
		wg.Add(1)
		go cajas[i].atenderClientes(&wg)
	}

	for i := 0; i < numClientes; i++ {
		cliente := fmt.Sprintf("Cliente%d", i+1)
		caja := cajaConColaMasCorta(cajas)
		caja.clientes <- cliente
	}

	for i := range cajas {
		close(cajas[i].clientes)
	}

	wg.Wait() // Espera a que todas las cajas terminen de atender

}
