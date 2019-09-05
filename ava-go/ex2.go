// 2. Escreva uma função que receba dois números inteiros e retorne dois números: menor, maior
package main

import "fmt"

func menorMaior(a, b int) (int, int) {
	var menor, maior int
	if a < b {
		menor = a
		maior = b
	} else {
		menor = b
		maior = a
	}
	return menor, maior
}


func main() {
    menor, maior := menorMaior(9, 50)
    fmt.Println(menor, maior)
}