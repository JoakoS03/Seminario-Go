package main

type Nodo struct {
	next *Nodo
	dato int
}

type List struct {
	head *Nodo
	size int
}

func (l List) New() List {
	return List{}
}

func (l List) IsEmpty(lista List) bool {
	if lista.size == 0 {
		return false
	}
	return true
}
func (l List) Len(lista List) int {
	return lista.size
}
func main() {

}
