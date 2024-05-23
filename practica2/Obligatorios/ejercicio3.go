package main

import (
	"fmt"
)

type Field struct {
	Value, Cant int
}

type OptimumSlice []Field

func New(s []int) (os OptimumSlice) {

	if len(s) == 0 {
		return
	}

	cant := 1
	var value int
	for i := 0; i < len(s)-1; i++ {
		value = s[i]
		if value == s[i+1] {
			cant++
		} else {
			os = append(os, Field{Value: value, Cant: cant})
			cant = 1
		}
	}

	os = append(os, Field{Value: s[len(s)-1], Cant: cant})

	return
}

func PrintOptimumSlice(os OptimumSlice) {
	for _, e := range os {
		fmt.Println("Value: ", e.Value, " Cant: ", e.Cant)
	}
}

func IsEmpty(os OptimumSlice) (ok bool) {
	if (os[0] == Field{}) {
		ok = true
	} else {
		ok = false
	}
	return
}

func Len(os OptimumSlice) (cant int) {
	for _, e := range os {
		cant += e.Cant
	}

	return
}

func main() {
	s := []int{2, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 7, 7, 9, 90, 90, 90, 100}
	os := New(s)
	fmt.Println(IsEmpty(os))
	fmt.Println(Len(os))
	//PrintOptimumSlice(os)
}
