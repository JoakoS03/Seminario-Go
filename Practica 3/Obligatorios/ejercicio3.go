package main

import "fmt"

/*
Defina una estructura Contact que contenga campos como Nombre,
Apellido, CorreoElectronico, y Telefono.
--
Cree una estructura llamada Agenda que contenga un mapa de Contact
con el correo electrónico como clave
*/

type Contact struct {
	Nombre, Apellido, CorreoElectronico string
	Telefono                            int
}

type Agenda map[string]Contact

/*
Implemente los siguientes métodos para la estructura Agenda:
i. AgregarContacto(contacto Contact): Agrega un nuevo
contacto a la agenda.

ii. EliminarContacto(correo string): Elimina un contacto de la
agenda dado su correo electrónico.

iii. BuscarContacto(correo string) Contact: Busca y devuelve
un contacto dado su correo electrónico.
*/
// Método para agregar un contacto a la agenda
func (a *Agenda) AgregarContacto(contacto Contact) {
	if (*a)[contacto.CorreoElectronico] == (Contact{}) {
		(*a)[contacto.CorreoElectronico] = contacto
	} else {
		err := fmt.Errorf("el contacto ya existe")
		fmt.Println("Error:", err)
	}
}

// Método para eliminar un contacto de la agenda
func (a *Agenda) EliminarContacto(correo string) {
	delete(*a, correo)
}

// Método para buscar un contacto en la agenda
func (a Agenda) BuscarContacto(correo string) Contact {
	if contacto, existe := a[correo]; existe {
		return contacto
	}
	return Contact{}
}

func main()
