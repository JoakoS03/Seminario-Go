package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	frase := "Los jueves juego al futbol"

	frase2 := replace(frase, "jueves")

	fmt.Println("La frase nueva es: " + frase2)
}

func replace(frase, jueves string) string {
	martesMin := []string{"m", "a", "r", "t", "e", "s"}
	martesMayus := []string{"M", "A", "R", "T", "E", "S"}
	frase2 := strings.ToLower(frase)

	var palabra string
	var index = 0

	fraseNueva := ""

	for i := 0; i < len(frase2); {
		index = strings.Index(frase2[i:], jueves)
		if index != -1 {
			palabra = frase2[index : index+6]
			if palabra == jueves {
				var martes string
				for j, letra := range frase[index : index+6] {

					if unicode.IsUpper(rune(letra)) {
						martes += martesMayus[j]
					} else {
						martes += martesMin[j]
					}
				}
				fraseNueva += frase[i:index]
				fraseNueva += martes
				i = index + 6
			}
		} else {
			fraseNueva += frase[i:]
			break
		}
	}

	return fraseNueva
}
