package main

import (
	"fmt"
	"strings"
)

type Date struct {
	day, month, year int
}

type Alumno struct {
	apellido, nombre, cOrigen, codigo string
	Fnacimiento                       Date
	titulo                            bool
}

type Nodo struct {
	sig  *Nodo
	dato Alumno
}

type List struct {
	head *Nodo
}

func (a Alumno) ToString(alumno Alumno) string {
	str := fmt.Sprintf("%s-%s-%s", alumno.nombre, alumno.apellido, alumno.codigo)
	return str
}

func (l List) New() List {
	return List{}
}

func (l List) Next(self List) List {
	if self.head != nil {
		return self
	}
	return List{}
}

func (l List) IsEmpty(lista List) bool {
	if lista.head == nil {
		return true
	} else {
		return false
	}
}

func (l List) PushFront(self *List, item Alumno) {
	var aux Nodo
	aux.dato = item
	aux.sig = self.head
	self.head = &aux
}

func PushBack(self List, item Alumno) *Nodo {
	var aux Nodo
	aux.dato = item
	if self.IsEmpty(self) {
		self.head = &aux
		return self.head
	} else {
		ult := self.head
		for ult.sig != nil {
			ult = ult.sig
		}
		ult.sig = &aux
		return ult
	}
}

func (l List) Len(lista List) int {
	cant := 0
	for lista.head != nil {
		cant++
		lista.head = lista.head.sig
	}
	return cant
}

func (l List) ToString(lista List) string {
	var str string
	current := l.head
	for current != nil {
		str += fmt.Sprintf("%s-", current.dato.ToString(current.dato))
		current = current.sig
	}
	return str
}

func (l List) FrontElement(lista List) Alumno {
	if lista.head != nil {
		return lista.head.dato
	}
	fmt.Println("Si ve este mensaje es que la lista esta vacia")
	return Alumno{}
}

func (l List) Remove(self List) int {
	return -1 //Preguntar
}
func MostrarNombre(name string) {
	fmt.Printf("El nombre es: %s", name)
}
func (l List) Iterate(self List, f func(string)) {
	for self.head.sig != nil {
		f(self.head.dato.nombre)
		self.head = self.head.sig
	}
}

func InformarBariloche(lista List) {
	for lista.head != nil {
		if strings.ToLower(lista.head.dato.cOrigen) == "bariloche" {
			fmt.Printf("Nombre : %s Apellido : %s - nacio en bariloche", lista.head.dato.nombre,
				lista.head.dato.apellido)
		}
		lista.head = lista.head.sig
	}
}

func AnioMasIngresantes(list List) int {
	anioMax := -1
	if list.IsEmpty(list) {
		return anioMax // Retorna -1 si la lista está vacía
	}
	cant := 0
	max := -1
	for list.head != nil {
		anioAct := list.head.dato.Fnacimiento.year
		temp := list.head
		for temp != nil {
			if temp.dato.Fnacimiento.year == anioAct {
				cant++
			}
			temp = temp.sig
		}
		if cant > max {
			max = cant
			anioMax = anioAct
		}
		cant = 0
		list.head = list.head.sig
	}
	return anioMax
}

func CarreraConMasInscriptos(list List) string {
	str := "Lista... vacia"
	if list.IsEmpty(list) {
		return str
	}
	cant := 0
	max := -1
	for list.head != nil {
		codAct := list.head.dato.codigo
		temp := list.head
		for temp != nil {
			if temp.dato.codigo == codAct {
				cant++
			}
			temp = temp.sig
		}
		if cant > max {
			max = cant
			str = codAct
		}
		cant = 0
		list.head = list.head.sig
	}
	return str
}

func main() {
	var lista List
	lista = lista.New()
	/*
		apellido, nombre, cOrigen, codigo string
		Fnacimiento                       time.Time
		titulo
	*/
	fmt.Println(lista.IsEmpty(lista))
	lista.PushFront(&lista, Alumno{
		apellido:    "sueyro",
		nombre:      "Joaquin",
		cOrigen:     "Berisso",
		codigo:      "LS",
		Fnacimiento: Date{day: 21, month: 9, year: 2003},
		titulo:      true,
	})

	lista.head = PushBack(lista, Alumno{
		apellido:    "sueyro",
		nombre:      "Pepe",
		cOrigen:     "Bariloche",
		codigo:      "LS",
		Fnacimiento: Date{day: 21, month: 9, year: 2003},
		titulo:      true,
	})

	InformarBariloche(lista)
	fmt.Println(AnioMasIngresantes(lista))
	fmt.Println(CarreraConMasInscriptos(lista))
	//EliminarIngresantes(lista)

	fmt.Println(lista.IsEmpty(lista))
	fmt.Println(lista.ToString(lista))
	fmt.Println(lista.Len(lista))
}
