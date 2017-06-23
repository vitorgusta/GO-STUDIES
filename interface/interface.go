package main

import (
	"fmt"
)

type Ser interface {
	Respirar() bool
}

type (
	Humano   struct{}
	Cachorro struct{}
	Pedra    struct{}
)

func (h Humano) Respirar() bool {
	return false
}

func (c Cachorro) Respirar() bool {
	return true
}

func (p Pedra) Respirar() bool {
	return false
}
func VerificandoSeEstaVivo(s Ser) {
	if s.Respirar() {
		fmt.Println("Está vivo.")
	} else {
		fmt.Println("Não está vivo.")
	}
}

func main() {
	humano := Humano{}
	cachorro := Cachorro{}
	pedra := Pedra{}

	VerificandoSeEstaVivo(humano)
	VerificandoSeEstaVivo(cachorro)
	VerificandoSeEstaVivo(pedra)
}
