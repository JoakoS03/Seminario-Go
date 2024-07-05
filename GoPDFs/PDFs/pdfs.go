/*
El nombre del paquete debe coincidir con el nombre de la carperta, el archivo
puedo no coincidir.
*/
package PDFs

import "fmt"

func Hola(nombre string) string {
	return fmt.Sprintf("Hola como estas: %s", nombre)
}
