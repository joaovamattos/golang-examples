// Faça um programa que utilize closures para criar uma
// função que gera (generator) os múltiplos de um número.
package main

import (
	"fmt"
	"math/rand"
)

func generator() func() int {
	n := 0
	var lista [60]int
	
	i := 0
    for i < 60 {
		lista[i] = i + 1
        i = i + 1
	}
	
	rand.Shuffle(len(lista), func(i, j int) { lista[i], lista[j] = lista[j], lista[i] })

	return func() int {
		n++
		return lista[n]
	}
}

func main() {
	sortear := generator()
	fmt.Println(sortear())
	fmt.Println(sortear())
	fmt.Println(sortear())
	fmt.Println(sortear())
	fmt.Println(sortear())
	fmt.Println(sortear())
}


