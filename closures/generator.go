// Faça um programa que utilize closures para criar uma
// função que gera (generator) os múltiplos de um número.
package main
import "fmt"

func main() {
	gerar := generator(2)
	fmt.Println(gerar())
	fmt.Println(gerar())
	fmt.Println(gerar())
	fmt.Println(gerar())
}

func generator(num int) func() int {
	n := num
	return func() int {
		n = n + n
		return n
	}
}