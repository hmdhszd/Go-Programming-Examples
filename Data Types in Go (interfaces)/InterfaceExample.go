package main

import "fmt"

// An interface allow a value to be
// of more than one type
//----------------------------------

type person struct {
	firstName string
	lastName  string
}

type SercetAgent struct {
	person
	LicenceToKill bool
}

//----------------------------------

type human interface {
	speak()
}

//----------------------------------

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

//----------------------------------

func (p person) speak() {
	fmt.Println("Person says: I am", p.firstName, p.lastName)
}

//----------------------------------

func (s SercetAgent) speak() {
	fmt.Println("SercetAgent says: I am", s.firstName, s.lastName)
}

//----------------------------------

func main() {

	sa1 := SercetAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		LicenceToKill: true,
	}

	sa2 := SercetAgent{
		person: person{
			firstName: "Leonard",
			lastName:  "Cohen",
		},
		LicenceToKill: false,
	}

	p1 := person{
		"hamid",
		"Hosseinzadeh",
	}

	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)


}
