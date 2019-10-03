//  3. Escreva um programa que demonstre a utilização de ponteiros em Go. Aborde ponteiros para inteiros, ponteiros para structs e a função new, no mínimo.
package main

import "fmt"

type discente struct {
    nome string
	idade int
	turma string
}

func NewAluno(nome string, idade int) *discente {
    a := discente{nome: nome, idade: idade}
    a.turma = "ADS"
    return &a
}

func main() {
	i := 42
	p := &i

	aluno := NewAluno("João", *p)
	alunop := &aluno
	
    fmt.Println(*alunop)
}