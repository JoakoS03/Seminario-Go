package main

import (
	"bytes"
	"fmt"
)

type Stack []int

// Crea una pila vacia
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

func (s Stack) ToString(pila Stack) string {
	var buffer bytes.Buffer
	if pila.isEmpty(pila) {
		return "La pila está vacía"
	} else {
		for _, elem := range pila {
			buffer.WriteString(fmt.Sprintf("%v ", elem))
		}
	}
	return buffer.String()
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

func (s Stack) Iterate(pila Stack) {
	for elem := range pila {
		fmt.Println(elem)
	}
}

func main() {
	var pila Stack
	fmt.Println(pila.New())
	fmt.Println(pila.isEmpty(pila))
	fmt.Println(pila.Len(pila))
	fmt.Println(pila.ToString(pila))
	fmt.Println(pila.FrontElement(pila))
	pila.Iterate(pila)

	fmt.Println("---------------------------")

	pila.Push(&pila, 12)
	fmt.Println(pila.Len(pila))
	fmt.Println(pila.isEmpty(pila))
	fmt.Println(pila.ToString(pila))
	fmt.Println(pila.FrontElement(pila))
}
