package main

import (
	"time"
	"crypto/sha256"
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
