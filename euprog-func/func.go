package main

import "fmt"

func main() {

	i, msg := soma(10, 2)
	if msg != "" {
		fmt.Print(i)
	}
	fmt.Print("REsult soma é", i)

}

func soma(n1, n2 int) (int, string) {
	result := n1 + n2
	if result > 10 {
		fmt.Print("Esse valor diosvfiudsfvi")
	}
	return result, ""

}
