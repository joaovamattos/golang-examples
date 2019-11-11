package main

import "fmt"

func trocar(n1, n2 *int) (int, int) {
	aux := *n1
	*n1 = *n2
	*n2 = aux
	return *n1, *n2
}

func elevar(x, n int) (int){
	if n == 0 {
		return 1
	}
	return x * elevar(x, n-1)	
}

func main() {
	n1 := 1
	n2 := 2
	fmt.Println(trocar(&n1, &n2))
	x := 2
	n := 3	
	fmt.Println(elevar(x, n))

}