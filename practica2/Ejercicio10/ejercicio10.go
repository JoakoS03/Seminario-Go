package main

import "fmt"

type Stack []int

//Crea una pila vacia
func (s Stack) New() Stack {
	return Stack{}
}
func (s Stack) isEmpty(pila Stack) bool {
	if pila == nil {
		return true
	}
	return false
}
func (s Stack) Len(pila Stack) int {
	len := 0
	if pila.isEmpty(pila) {
		return 0
	} else {
		for _, i := range pila {
			len++
			i++
		}
	}
	return len
}

func (s Stack) ToStirng(pila Stack) string {
	str := ""
	if pila.isEmpty(pila) {
		return "La pila esta vacia"
	} else {
		for elem := range pila {
			str += string(elem)
		}
	}
	return str
}
func (s Stack) FrontElement(pila Stack) int {
	if pila.isEmpty(pila) {
		//Quiere decir que la pila esta vacia
		return -1
	}
	return pila[0]
}

func (s Stack) Push(pila *Stack, elem int) {
	*pila = append(*pila, elem)
}
func (s Stack) Pop(pila Stack) int {
	if pila.isEmpty(pila) {
		return -1
	}
	return pila[len(pila)-1]
}

func (s Stack) Iterate(pila Stack, f func(int) int) {
	for elem := range pila {
		fmt.Println(elem)
		f(elem)
	}
}

func main() {
	var pila Stack
	fmt.Println(pila.New())
	fmt.Println(pila.isEmpty(pila))
	fmt.Println(pila.Len(pila))
	fmt.Println(pila.ToStirng(pila))
	fmt.Println(pila.FrontElement(pila))
	fmt.Println(pila.Iterate(pila, Sumar))
}
