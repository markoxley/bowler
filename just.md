package main

import "fmt"

type Animal interface {
	Legs() int
}

type Dog struct {
	legs int
}

func (d Dog) Legs() int {
	return 4
}

func create[T Animal]() *T {
	t := new(T)
	return t
}

func createList[T Animal]() []*T {
	a := make([]*T, 5)
	for i := range a {
		a[i] = new(T)
	}
	return a
}
func main() {
	d := create[Dog]()
	fmt.Printf("%T\t%v\n", d, d)
	fmt.Println(d.Legs())

	a := createList[Dog]()
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Println(len(a))
}
