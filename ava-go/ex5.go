// 5. Escreva uma função que leia um número não determinado de valores inteiros positivos e retorne a média entre eles
package main

import "fmt"

func media(nums ...int) {
	total := 0
	i:= 0
    for _, num := range nums {
		total += num
		i += 1
    }
    fmt.Println(total/i)
}

func main() {
    nums := []int{2, 3, 4}
    media(nums...)
}