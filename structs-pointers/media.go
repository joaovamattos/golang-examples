// 2. Escreva um programa que crie um tipo Aluno, que possua, além de nome, nota1, nota2 e nota3. Escreva uma função que retorne a média de um aluno.
package main

import "fmt"

type aluno struct {
   nome string
   nota1, nota2, nota3 int
}

func (a aluno) media() int {
	return (a.nota1 + a.nota2 + a.nota3)/3
}

func main() {

    a := aluno{nome: "João", nota1: 90, nota2: 90, nota3: 80}
    fmt.Println("media:", a.media())
}