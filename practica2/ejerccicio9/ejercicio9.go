package main

import "fmt"

type Nodo struct {
	sig  *Nodo
	dato int
}

type List struct {
	head *Nodo
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
		return false
	} else {
		return true
	}
}

func (l List) PushFront(self *List, item int) {
	var aux Nodo
	aux.dato = item
	aux.sig = self.head
	self.head = &aux
}

func (l List) PushBack(self List, item int) {
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
	for lista.head.sig != nil {
		cant++
		lista.head = lista.head.sig
	}
	return cant
}

func (l List) ToString(lista List) string {
	var str string
	for lista.head.sig != nil {
		str += string(lista.head.dato)
		lista.head = lista.head.sig
	}
	return str
}

func (l List) FrontElement(lista List) int {
	if lista.head.sig != nil {
		return lista.head.dato
	}
	fmt.Println("Si ve este mensaje es que la lista esta vacia")
	return -1
}

func (l List) Remove(self List) int {
	return -1 //Preguntar
}
func SumarValores(num int) int {
	return num + num
}
func (l List) Iterate(self List, f func(int)) int {
	suma := 0
	for self.head.sig != nil {
		suma += f(self.head.dato)
		self.head = self.head.sig
	}
	return suma
}
func main() {

}
