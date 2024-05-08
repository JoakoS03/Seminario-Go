package main

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
	var aux Nodo
	ult := self.head
	aux.dato = item
	for ult.sig != nil {
		ult = ult.sig
	}
	ult.sig = &aux
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
	return -1
}
