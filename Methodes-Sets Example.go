package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello from", p.name)
}

type human interface {
	speak()
}

func saySomeThing(h human) {
	h.speak()
}

func main() {

	p1 := person{
		name: "Hamid",
	}

	saySomeThing(&p1)

	p1.speak()

}
