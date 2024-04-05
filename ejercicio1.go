package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	frase := "Los jueves juego al futbol y los JUEVES al tenis"

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
	i := 0
	for i < len(frase2) {
		index = strings.Index(frase2[i:], jueves)
		if index != -1 {
			palabra = frase2[i+index : i+index+6]
			if palabra == jueves {
				var martes string
				for j, letra := range frase[i+index : i+index+6] {

					if unicode.IsUpper(rune(letra)) {
						martes += martesMayus[j]
					} else {
						martes += martesMin[j]
					}
				}
				fraseNueva += frase[i : i+index]
				fraseNueva += martes
				i = i + index + 6
			}
		} else {
			fraseNueva += frase[i:]
			break
		}
		fmt.Println(i)
	}

	return fraseNueva
}
