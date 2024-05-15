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

func (alumno Alumno) ToString() string {
	str := fmt.Sprintf("%s-%s-%s", alumno.nombre, alumno.apellido, alumno.codigo)
	return str
}

func New() List {
	return List{}
}

func IsEmpty(lista List) bool {
	if lista.head == nil {
		return true
	} else {
		return false
	}
}

func PushFront(self *List, item Alumno) {
	var aux Nodo
	aux.dato = item
	aux.sig = self.head
	self.head = &aux
}

func PushBack(self *List, item Alumno) {
	var aux Nodo
	aux.dato = item
	if IsEmpty(*self) {
		self.head = &aux
	} else {
		ult := self.head
		for ult.sig != nil {
			ult = ult.sig
		}
		ult.sig = &aux

	}
}

func Len(lista List) int {
	cant := 0
	for lista.head != nil {
		cant++
		lista.head = lista.head.sig
	}
	return cant
}

func ToString(l List) string {
	var str string
	current := l.head
	for current != nil {
		str += fmt.Sprintf("%s-", current.dato.ToString())
		current = current.sig
	}
	return str
}

func FrontElement(lista List) Alumno {
	if lista.head != nil {
		return lista.head.dato
	}
	fmt.Println("Si ve este mensaje es que la lista esta vacia")
	return Alumno{}
}

func Remove(self *List, alumno Alumno) {
	var ant *Nodo
	actual := self.head

	for actual != nil {
		if actual.dato == alumno {
			if ant == nil {
				self.head = actual.sig
			} else {
				ant.sig = actual.sig
			}
			break
		}
		ant = actual
		actual = actual.sig
	}
}
func MostrarNombre(name string) {
	fmt.Printf("El nombre es: %s", name)
}

/*func Iterate(self List, f func(string)) {
	for self.head.sig != nil {
		f(self.head.dato.nombre)
		self.head = self.head.sig
	}
}*/

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
	if IsEmpty(list) {
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
	if IsEmpty(list) {
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

func EliminarIngresantes(lista *List) {
	actual := lista.head
	for actual != nil {
		siguiente := actual.sig // Guardar referencia al siguiente nodo antes de modificar actual
		if actual.dato.titulo == false {
			Remove(lista, actual.dato)
		}
		actual = siguiente // Mover actual al siguiente nodo
	}
}

func main() {
	var lista List
	lista = New()
	/*
		apellido, nombre, cOrigen, codigo string
		Fnacimiento                       time.Time
		titulo
	*/
	fmt.Println(IsEmpty(lista))
	PushFront(&lista, Alumno{
		apellido:    "sueyro",
		nombre:      "Joaquin",
		cOrigen:     "Berisso",
		codigo:      "LS",
		Fnacimiento: Date{day: 21, month: 9, year: 2003},
		titulo:      true,
	})

	PushBack(&lista, Alumno{
		apellido:    "sueyro",
		nombre:      "Pepe",
		cOrigen:     "Bariloche",
		codigo:      "LS",
		Fnacimiento: Date{day: 21, month: 9, year: 2003},
		titulo:      false,
	})

	PushBack(&lista, Alumno{
		apellido:    "Pepe",
		nombre:      "Mujica",
		cOrigen:     "Bariloche",
		codigo:      "LS",
		Fnacimiento: Date{day: 21, month: 9, year: 2003},
		titulo:      true,
	})

	fmt.Println(ToString(lista))
	fmt.Println("--------------------------------------------")
	InformarBariloche(lista)
	fmt.Println("--------------------------------------------")
	fmt.Println(AnioMasIngresantes(lista))
	fmt.Println("--------------------------------------------")
	fmt.Println(CarreraConMasInscriptos(lista))
	fmt.Println("--------------------------------------------")
	EliminarIngresantes(&lista)
	fmt.Println("--------------------------------------------")

	fmt.Println(IsEmpty(lista))
	fmt.Println("--------------------------------------------")
	fmt.Println(ToString(lista))
	fmt.Println("--------------------------------------------")
	fmt.Println(Len(lista))
	fmt.Println("--------------------------------------------")
}
