package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"
)

func Primo(numero int, list *[]int) {
	if numero < 2 {
		return
	}
	if numero == 2 {
		*list = append(*list, numero)
		return
	}
	if numero%2 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(numero))); i += 2 {
		if numero%i == 0 {
			return
		}
	}
	*list = append(*list, numero)
}

func main() {
	// Definir un parámetro de línea de comando para un número entero positivo
	numero := flag.Int("n", 0, "un numero entero positivo")

	flag.Parse()

	// Verificar que se haya ingresado un número positivo
	if *numero <= 0 {
		log.Fatal("Debes ingresar un número entero positivo")
	}

	listPrimos := []int{}
	for i := 0; i <= *numero; i++ {
		go Primo(i, &listPrimos)
	}

	// Esperar un tiempo arbitrario para permitir que las gorutinas terminen
	time.Sleep(2 * time.Second)

	fmt.Println("Números primos:", listPrimos)
	fmt.Printf("El número ingresado es: %d\n", *numero)
}
