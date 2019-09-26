// Faça um programa que utilize closures para gerar
// sequências de 6 números pseudo-aleatórios não repetidos
// (a cada 6 chamadas a sequência é reinicializada).
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println(randInt(1, 60))
}

func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func generator(num int) func() int {
	n := num
	return func() int {
		n = n + n
		return n
	}
}