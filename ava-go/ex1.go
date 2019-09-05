// 1. Escreva uma função que receba dois números inteiros e retorne o menor entre eles. 
package main

import "fmt"

func menor(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}


func main() {
    res := menor(9, 2)
    fmt.Println(res)
}
