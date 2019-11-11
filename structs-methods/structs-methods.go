// Crie uma estrutura ponto, com os valores X e Y e escreva um método distance que recebe por parâmetro uma variável da estrutura ponto e calcula a distância entre estes dois pontos.
package main

import (
	"fmt"
	"math"
)

type Ponto struct {
    x int
    y int
}

func (p Ponto) distance(outroPonto Ponto) float64 {

	var cat1, cat2 int

	if p.x > outroPonto.x {
		cat1 = p.x - outroPonto.x
	} else {
		cat1 = outroPonto.x - p.x 
	}

	if p.y > outroPonto.y {
		cat2 = p.y - outroPonto.y
	} else {
		cat2 = outroPonto.y - p.y 
	}

	hip := cat1 * cat1 + cat2 * cat2
	dist := math.Sqrt(float64(hip)) 
	return dist

}
func main() {
	p1 := Ponto{x: 2, y: 2}
	p2 := Ponto{x: 3, y: 3}
	fmt.Println(p1.distance(p2))
}