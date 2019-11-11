package main

import "fmt"

type dataStruct interface {
    add(i int64)
	remove()
	imprimir()
}

type pilha struct {
    stack []int64
}
type fila struct {
    queue []int64
}

func (p *pilha) add(i int64) {
    p.stack = append(p.stack, i)
}
func (p *pilha) remove() {
    p.stack = p.stack[:len(p.stack) - 1]
}

func (p *pilha) imprimir(){
	fmt.Println(p.stack)
}

func (f *fila) add(i int64) {
    f.queue = append(f.queue, i)
}
func (f *fila) remove() {
	f.queue = append(f.queue[:0], f.queue[1:]...)
}

func (f *fila) imprimir(){
	fmt.Println(f.queue)
}

func measure(d dataStruct) {
	d.add(1)
	d.imprimir()
	d.add(2)
	d.add(3)
	d.imprimir()
	d.remove()
	d.remove()
	d.imprimir()
}

func main() {
	fmt.Println("Pilha")
    s := pilha{}
	measure(&s)

	fmt.Println("Fila")
    q := fila{}
    measure(&q)
}