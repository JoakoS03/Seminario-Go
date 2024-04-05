package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Tengo que hacer una funcion que reciba una palabra y
en la funcion pida ingresar un texto y si la palabra pasada en el
parametro se encuentra, invertir las mayusculas y minusculas dependiendo de como
se ingreso la palabra
*/
func main() {
	var word string
	//fmt.Println("Ingrese un palabra: ")
	//fmt.Scanln(&word)
	word = "helado"
	fraseNueva := cambiarPalabra(word)

	println("La frase nueva es: \n" + fraseNueva)
}

func cambiarPalabra(word string) (fraseNueva string) {
	var frase string
	//println("Ingrese una frase")
	//fmt.Scan(&frase)
	frase = "me gusta el helado con crema y el HELADO con chocolate"
	frase2 := strings.ToLower(frase)
	i := 0
	for i < len(frase2) {
		index := strings.Index(frase2[i:], strings.ToLower(word))

		if index != -1 {
			palabra := frase[i+index : i+index+len(word)]
			var wordNew string
			for _, letra := range palabra {
				wordNew += string(intercambiar(letra))
			}
			fraseNueva += frase[i : i+index]
			fraseNueva += wordNew
			i = i + index + len(word)
		} else {
			fraseNueva += frase[i:]
			i = len(frase2)
		}
		fmt.Println(fraseNueva)
		println(i)
	}
	return
}
func intercambiar(letra rune) (word string) {
	if unicode.IsUpper(rune(letra)) {
		word += string(unicode.ToLower(letra))
	} else {
		word += string(unicode.ToUpper(letra))
	}
	return
}
