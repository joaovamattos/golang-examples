// 3. Escreva uma função que receba um número arbitrário de números e retorne a soma de todos eles. 
package main

import "fmt"

func sum(nums ...int) {
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}