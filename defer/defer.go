package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Print("Fim da função main\r\n")
	}()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\r\n", i)
		fmt.Printf("Dentro do for %d\r\n", i)
	}
}
