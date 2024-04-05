package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var frase string

	fmt.Scanf("Ingrese J o M para remplazar martes por jueves y miercoles por automovil", &tipo)

	println("Ingrese un frase: ")

	fmt.Scan(&frase)
	word := 0
	if tipo == "J" {
		word := 7
	} else {
		word := 9
	}

	frase2 := replace(frase, word)

	fmt.Println("La frase nueva es: " + frase2)
}

func replace(frase string, lenght int) string {
	var wordRepacle string

	frase2 := strings.ToLower(frase)
	if lenght == 7 {
		wordRepacle = "jueves"
	} else {
		wordRepacle = "miercoles"
	}
	var palabra string
	var index = 0
	fraseNueva := ""

	for i := 0; i < len(frase2); {
		index = strings.Index(frase2[i:], wordRepacle)
		if index != -1 {

			palabra = frase2[i+index : i+index+lenght]
			fmt.Println(palabra, wordRepacle)
			if palabra == wordRepacle {
				var palabraReplace string

				for j, letra := range frase[i+index : i+index+lenght] {
					if lenght == 7 {
						palabraReplace += replaceJueves(j, rune(letra))
					} else {
						palabraReplace += replaceMiercoles(j, rune(letra))
					}
				}

				fraseNueva += frase[i : i+index]
				fraseNueva += palabraReplace
				i = i + index + lenght
			}
		} else {
			fraseNueva += frase[i:]
			i = len(frase2)
		}
	}

	return fraseNueva
}

func replaceJueves(indice int, letra rune) string {
	martesMin := []string{"m", "a", "r", "t", "e", "s"}
	martesMayus := []string{"M", "A", "R", "T", "E", "S"}
	var palabraReplace string

	if unicode.IsUpper(rune(letra)) {
		palabraReplace += martesMayus[indice]
	} else {
		palabraReplace += martesMin[indice]
	}
	return palabraReplace
}

func replaceMiercoles(j int, letra rune) string {
	autoMin := strings.Split("automovil", "")
	autoMayus := strings.Split("AUTOMOVIL", "")

	var palabraReplace string

	if unicode.IsUpper(letra) {
		palabraReplace += autoMayus[j]
	} else {
		palabraReplace += autoMin[j]
	}

	return palabraReplace
}
