package main

import (
	"errors"
	"log"
)

var ErrPanettoneDeChocolate = errors.New("Panettone de chocolate é miado")

var ErrPanettoneSeco = errors.New("Panettone ta seco")

type Origem int

const (
	Esposa Origem = iota
	Sogra
	Cunhada
	Voce
)

type panettone struct {
	TemChocolate bool
	TaSeco       bool
	VeioDe       Origem
}

func (p *panettone) verificaPanettone() (err error) {
	if p.TemChocolate {
		err = ErrPanettoneDeChocolate
		return
	}
	if p.TaSeco {
		err = ErrPanettoneSeco
		return
	}
	return
}
func comentarioSarcastico() {
	log.Println("iniciando comentario sarcastico")
}

func resmungaBaixinho() {
	log.Println("iniciando cresmungos")
}

func main() {
	p := panettone{
		TemChocolate: false,
		TaSeco:       true,
		VeioDe:       Sogra,
	}

	err := p.verificaPanettone()
	if err != nil {
		log.Println(err)
		if err == ErrPanettoneDeChocolate {
			panic(err)
		}
		if err == ErrPanettoneSeco {
			if p.VeioDe == Sogra ||
				p.VeioDe == Cunhada {
				comentarioSarcastico()
			} else if p.VeioDe == Esposa {
				resmungaBaixinho()
			}
		}
	}
}
