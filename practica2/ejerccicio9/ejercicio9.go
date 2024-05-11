package main

import (
	"fmt"
	"time"
)

type Alumno struct {
	apellido, nombre, cOrigen, codigo string
	Fnacimiento                       time.Time
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
	return fmt.Sprintf("%v", alumno.nombre, alumno.apellido, alumno.Fnacimiento, alumno.codigo)
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

func (l List) PushBack(self List, item Alumno) {
	if self.IsEmpty(self) {
		self.PushFront(&self, item)
	} else {
		var aux Nodo
		ult := self.head
		aux.dato = item
		for ult.sig != nil {
			ult = ult.sig
		}
		ult.sig = &aux
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
		str += fmt.Sprintf("%v-", current.dato.ToString)
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
func SumarValores(num int) int {
	return num + num
}
func (l List) Iterate(self List, f func(int) int) {
	suma := 0
	for self.head.sig != nil {
		suma += f(self.head.dato.nombre)
		self.head = self.head.sig
	}
}

func main() {
	var lista List
	lista = lista.New()

	fmt.Println(lista.IsEmpty(lista))
	fmt.Println(lista.IsEmpty(lista))
	fmt.Println(lista.ToString(lista))
	fmt.Println(lista.Len(lista))
}
