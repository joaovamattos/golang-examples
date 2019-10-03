// 1. Escreva um programa para criar novos tipos de dados conforme abaixo:
// - Horário: composto de hora, minutos e segundos
// - Data: composto de dia, mês e ano
// - Compromisso: composto de data, horário e descrição
package main

import "fmt"

type horario struct {
    hora int
    minuto int
    segundo int
}

type data struct {
    dia int
    mes string
    ano int
}

type compromisso struct {
    d data
    h horario
	descricao string
}

func main() {

    h := horario{hora: 19, minuto: 45, segundo: 37}
    d := data{dia: 20, mes: "janeiro", ano: 2019}
	c := compromisso{d: d, h: h, descricao: "Festa" }

    fmt.Println(c)
}