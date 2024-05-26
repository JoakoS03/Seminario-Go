package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Transaccion struct {
	Monto       float64
	Id          string
	IdRecividor string
	Timestamp   time.Time
}

type Block struct {
	PreviusHash string
	Hash        string
	Data        Transaccion
	Timestamp   time.Time
}

type Wallet struct {
	Id       string
	Nombre   string
	Apellido string
}

type Blockchain []Block

func crearRandom() string {
	return strconv.Itoa(rand.Intn(100000))
}
func CreateWallet(nombre, apellido string) Wallet {
	random := crearRandom()
	return Wallet{Id: random, Nombre: nombre, Apellido: apellido}
}

func CrearTransaccion(monto float64, id, idRecividor string) Transaccion {
	return Transaccion{Monto: monto, Id: id, IdRecividor: idRecividor, Timestamp: time.Now()}
}

func CrearBlock(transaccion Transaccion, previusHash string) Block {
	timestamp := time.Now()
	dataToHash := fmt.Sprintf("%v%v%v", previusHash, transaccion, timestamp)
	hash := sha256.Sum224([]byte(dataToHash))
	hashString := hex.EncodeToString(hash[:])

	return Block{
		Hash:        hashString,
		PreviusHash: previusHash,
		Timestamp:   timestamp,
		Data:        transaccion,
	}
}

func VerificarConsistencia(cadena Blockchain) bool {
	consistente := true
	i := 0
	for consistente && i < len(cadena)-1 {
		if cadena[i+1].PreviusHash != cadena[i].Hash {
			consistente = false
		}
		i++
	}

	return consistente
}

func SaldoUsuario(id string, cadena Blockchain) (suma float64) {
	suma = 0
	for i := 0; i < len(cadena); i++ {
		if cadena[i].Data.IdRecividor == id {
			suma += cadena[i].Data.Monto
		}
	}
	return
}

func main() {
	var bc Blockchain
	user1 := CreateWallet("Joaquin", "Sueyro")
	user2 := CreateWallet("Luis", "Sueyro")
	transaccion1 := CrearTransaccion(12222, user1.Id, user2.Id)

	user3 := CreateWallet("Ricardo", "Perez")
	user4 := CreateWallet("Alfonso", "Preira")
	t2 := CrearTransaccion(1258, user3.Id, user4.Id)

	user5 := CreateWallet("Javier", "Perez")
	user6 := CreateWallet("Cristina", "Perez")
	t3 := CrearTransaccion(5000, user5.Id, user6.Id)

	t4 := CrearTransaccion(6000, user3.Id, user2.Id)

	block1 := CrearBlock(transaccion1, "00000000") //Bloque inical.
	block2 := CrearBlock(t2, block1.Hash)
	block3 := CrearBlock(t3, block2.Hash)
	block4 := CrearBlock(t4, block3.Hash)

	bc = append(bc, block1)
	bc = append(bc, block2)
	bc = append(bc, block3)
	bc = append(bc, block4)

	fmt.Println("------------Punto 5---------------")
	fmt.Println(VerificarConsistencia(bc))
	fmt.Println("------------Punto 4---------------")

	saldo := SaldoUsuario(user2.Id, bc)
	fmt.Println("El saldo del usuario: ", user2.Id, " es ", saldo)

	/*
		fmt.Println(bc[0].Hash)
		fmt.Println(bc[1].PreviusHash)

		fmt.Println("---------------------------")

		/*
		fmt.Println(bc[1].Hash)
		fmt.Println(bc[2].PreviusHash)*/
}
