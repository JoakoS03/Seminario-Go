package main
/*
Se puede remplazar el flag.String por el flag.Int para ingresar un valor numero pero el string puedo ver que pasa si ingreso un valor negativo
*/
import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {

	//Declas el parametro por linea de comando a traves del comando -number
	num := flag.String("number", "", " ")

	//Parsea los args de linea de comandos
	flag.Parse()

	//Verifica que se ingreso algo
	if *num == "" {
		log.Fatal("Debes ingresar un numero forro")
	}

	n, err := strconv.Atoi(*num)
	if err != nil || n <= 0 {
		log.Fatal("Tenes que ingresar un numero entero valido")
	}

	fmt.Printf("El número ingresado es: %d\n", n)
}
