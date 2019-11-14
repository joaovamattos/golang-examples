package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
    arg  float64
    prob string
}
func (e *ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("%s: %f", e.prob, e.arg)
}

func Sqtr(x float64) (float64, error){
	if(x < 0){
		return -1, &ErrNegativeSqrt{arg: x, prob: "Impossível obter a raiz de número negativo"}
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqtr(2))
	fmt.Println(Sqtr(-2))
}