// 4. Escreva uma função que receba um número arbitrário de números e retorne dois números: maior ,menor
package main

import "fmt"

func maiorMenor(nums ...int) (int, int){
	maior := 0
	menor := nums[1]
    for _, num := range nums {
		if num > maior{
			maior = num
		} 
		if num < menor{
			menor = num
		}
		
    }
	return maior, menor
}


func main() {
    nums := []int{1, 2, 3, 4, 99, 0}
    maior, menor := maiorMenor(nums...)
    fmt.Println(maior, menor)
}