package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var frase, word1, word2 string
	println("Ingrese un frase: ")

	fmt.Scan(&frase)

	fmt.Println("Ingrese la palabra que quiere reemplazar")
	fmt.Scan(&word1)

	fmt.Println("Ingrese la palabra por la que quiere reemplazar que tengan la misma cantidad de letras")
	fmt.Scan(&word2)

	frase2 := replace(frase, word1, word2)

	fmt.Println("La frase nueva es: " + frase2)
}

func replace(frase, word1, word2 string) string {

	frase2 := strings.ToLower(frase)

	var palabra string
	var index = 0
	fraseNueva := ""

	for i := 0; i < len(frase2); {

		index = strings.Index(frase2[i:], word1)
		if index != -1 {

			palabra = frase2[index : index+len(word1)]

			if palabra == word1 {
				var palabraReplace string

				for j, letra := range frase[index : index+len(word1)] {
					palabraReplace += replaceWord(j, rune(letra), word2)
				}

				fraseNueva += frase[i:index]
				fraseNueva += palabraReplace
				i = index + len(word1)
			}
		} else {
			fraseNueva += frase[i:]
			i = len(frase2)
		}
	}

	return fraseNueva
}

func replaceWord(j int, letra rune, word string) string {
	Min := strings.Split(strings.ToLower(word), "")
	Mayus := strings.Split(strings.ToUpper(word), "")

	var palabraReplace string

	if unicode.IsUpper(letra) {
		palabraReplace += Mayus[j]
	} else {
		palabraReplace += Min[j]
	}

	return palabraReplace
}
