package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sync"
)

func definirParametro() *int {
	num := flag.Int("n", 0, "Numero entero positivo")
	flag.Parse()

	fmt.Printf("Numero: %d\n", *num)
	if *num <= 0 {
		log.Fatal("Debe ingresar un numero valido")
	}

	return num
}

func Primo(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func buscarPrimos(rango []int, primos *[]int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	var localPrimos []int

	for i := rango[0]; i <= rango[1]; i++ {
		if Primo(i) {
			localPrimos = append(localPrimos, i)
		}
	}
	for _, p := range localPrimos {
		*primos = append(*primos, p)
	}
	return *primos
}

func main() {
	num := definirParametro()
	veces_a_dividir := 4
	//Numero primos
	primos := []int{}

	aux := *num / veces_a_dividir

	//Rango a recorrer por la go rutinas

	var wg sync.WaitGroup

	for cant := 0; cant < veces_a_dividir; cant++ {
		rango := make([]int, 2)
		rango[0] = cant * aux
		rango[1] = (cant+1)*aux - 1

		wg.Add(1)
		go buscarPrimos(rango, &primos, &wg)
	}

	wg.Wait()
	fmt.Printf("Numeros primos: %v", primos)

}
