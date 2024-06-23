package main

import (
	"fmt"
	"sync"
	"time"
)

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

type Agenda struct {
	contactos map[string]Contact
	mutex     sync.Mutex
}

// Crea una nueva agenda
func NuevaAgenda() Agenda {
	return Agenda{
		contactos: make(map[string]Contact),
	}
}

// Agrega un contacto nuevo
func (a *Agenda) AgregarContacto(contacto Contact) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.contactos[contacto.CorreoElectronico] == (Contact{}) {
		a.contactos[contacto.CorreoElectronico] = contacto
	} else {
		err := fmt.Errorf("el contacto ya existe")
		fmt.Println("Error:", err)
	}
}

// Método para eliminar un contacto de la agenda
func (a *Agenda) EliminarContacto(correo string) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	delete(a.contactos, correo)
}

// Método para buscar un contacto en la agenda
func (a *Agenda) BuscarContacto(correo string) Contact {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if contacto, existe := a.contactos[correo]; existe {
		return contacto
	}

	return Contact{}

}

func (a *Agenda) imprimirContactos() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	for k, v := range a.contactos {
		fmt.Println("Clave:", k, "Dato:", v)
	}

}

func main() {
	a := NuevaAgenda()

	go a.AgregarContacto(Contact{"Rodrigo", "Mora", "r@gmail.com", 123})
	go a.AgregarContacto(Contact{"Lionel", "Messi", "m@gmail.com", 1234})
	go a.AgregarContacto(Contact{"Gonzalo", "Montiel", "gm@gmail.com", 12345})
	go a.AgregarContacto(Contact{"Enzo", "Fernandez", "f@gmail.com", 123456})
	go a.AgregarContacto(Contact{"Julian", "Alvarez", "ja@gmail.com", 1234567})

	
	time.Sleep(1 * time.Second)

	a.imprimirContactos()

	contacto := a.BuscarContacto("m@gmail.com")
	time.Sleep(1 * time.Second)
	fmt.Println("Resultado de búsqueda:", contacto)

	go a.EliminarContacto("m@gmail.com")

	
	time.Sleep(1 * time.Second)

	a.imprimirContactos()

}
