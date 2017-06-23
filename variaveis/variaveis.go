package main

import (
	"fmt"
)

var ( //declarndo multiplas variaveis
	nome                   = "Cezar"
	idade                  = 42
	url, path, query, page = "example.org", "/search", "&q=golang", 1
)

const pi = 3.141592

func main() {
	var x, y int = 1, 2

	var s string
	var a string

	a = "texto 1"
	b := "texto 2"

	ola := func() {
		fmt.Printf("olá função anonima!\r\n")
	}

	fmt.Printf("a tipo: %T\r\n", a)
	fmt.Printf("b tipo: %T\r\n", b)
	fmt.Printf("pi tipo: %T\r\n", pi)
	fmt.Printf("ola tipo: %T\r\n", ola)

	fmt.Printf("valor de a = %v\r\n", a)
	fmt.Printf("valor de b = %v\r\n", b)
	fmt.Printf("valor de pi = %v\r\n", pi)

	fmt.Printf("valor de x = %v\r\n", x)
	fmt.Printf("valor de y = %v\r\n", y)

	fmt.Printf("valor de s = %q\r\n", s)

	ola()

}
