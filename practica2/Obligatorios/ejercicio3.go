package main

import (
	"fmt"
)

type Field struct {
	Value, Cant int
}

type OptimumSlice []Field

func PrintOptimumSlice(os OptimumSlice) {
	for _, e := range os {
		fmt.Println("Value: ", e.Value, " Cant: ", e.Cant)
	}
}

func PrintSlice(s []int) {
	for _, e := range s {
		fmt.Println("Value: ", e)
	}
}

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

func LastElement(os OptimumSlice) (lastElement Field) {
	if IsEmpty(os) {
		return
	}
	lastElement = os[len(os)-1]
	/*for i := range os {
		lastElement = os[i]
	}*/
	return
}

func FrontElement(os OptimumSlice) (frontElement Field) {
	if IsEmpty(os) {
		return
	}
	frontElement = os[0]
	return
}

func SliceArray(os OptimumSlice) (slice []int) {
	for _, elemOs := range os {
		for j := 0; j < elemOs.Cant; j++ {
			slice = append(slice, elemOs.Value)
		}
	}
	return
}
func Insert(os OptimumSlice, element, position int) (newElements OptimumSlice) {

	if position < 0 || position > Len(os) {
		fmt.Println("Index out of range")
		return
	}

	newElements = OptimumSlice{}
	index := 0

	for _, field := range os {
		if position > index && position <= index+field.Cant {
			if position == index+field.Cant {
				// Lo mete al final
				newElements = append(newElements, field)
				newElements = append(newElements, Field{element, 1})
			} else {
				// Corta el campo y lo mete en el medio
				newElements = append(newElements, Field{field.Value, position - index})
				newElements = append(newElements, Field{element, 1})
				newElements = append(newElements, Field{field.Value, field.Cant - (position - index)})
			}
		} else {
			newElements = append(newElements, field)
		}
		index += field.Cant
	}

	// If the position is at the very end, append the new elements
	if position == Len(os) {
		newElements = append(newElements, Field{element, 2})
	}

	return
}
func main() {
	s := []int{2, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 7, 7, 9, 90, 90, 90, 100}
	os := New(s)

	fmt.Println(IsEmpty(os))
	fmt.Println(Len(os))

	fe := FrontElement(os)
	le := LastElement(os)
	fmt.Println(fe)
	fmt.Println(le)

	sn := SliceArray(os)
	fmt.Println(sn)

	os = Insert(os, 50, 5)

	fmt.Println(os)
	//PrintSlice(sn)
	//PrintOptimumSlice(os)
}
