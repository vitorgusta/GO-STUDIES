package main

import (
	"encoding/json"
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
}

type Localidade struct {
	X     int    `json:"valor_x"`
	Y     int    `json:"valor_y"`
	Nome  string `json:"nome_da_localidade"`
	valor int    // campo que inicia com letra minuscula é "private"
}

func (p *Localidade) soma(l Localidade) {
	p.X += l.X
	p.Y += l.Y
}

func main() {

	aluno := Aluno{
		Nome:  "Cesar",
		Idade: 42,
	}
	fmt.Println("Aluno:", aluno)

	minhaCasa := Localidade{3, 4, "casa", 500}

	estruturaVazia := Aluno{}

	escolaLocalidade := Localidade{
		Y:    100,
		X:    200,
		Nome: "escola",
	}
	var outraLocalidade Localidade
	outraLocalidade.X = 10
	outraLocalidade.Y = 20
	outraLocalidade.Nome = "trabalho"

	fmt.Println("Minha casa:", minhaCasa)
	fmt.Println("outra localidade:", outraLocalidade)
	fmt.Println("escola:", escolaLocalidade)
	fmt.Println("aluno:", aluno)
	fmt.Println("estrutura vazia: %q\r\n", estruturaVazia)

	minhaCasa.soma(outraLocalidade)

	fmt.Println("localidade minha casa somada com outra localidade:", minhaCasa)

	fmt.Println("Minha casa: %v\r\n", minhaCasa)
	fmt.Println("Minha casa: %+v\r\n", minhaCasa)

	j, err := json.Marshal(minhaCasa)
	if err != nil {
		panic(err)
	}

	fmt.Println("Minha casa json:", string(j))

	novaLocalidade := Localidade{}
	err = json.Unmarshal(j, &novaLocalidade)
	if err != nil {
		panic(err)
	}

	fmt.Println("Pondo depois do unmarshal:", minhaCasa)

}
