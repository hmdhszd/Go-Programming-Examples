package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {

	p1 := person{
		Firstname: "James",
		Lastname:  "Bond",
		Age:       99,
	}

	p2 := person{
		Firstname: "Hamid",
		Lastname:  "Hosseinzadeh",
		Age:       30,
	}

	//-------Marshal--------------------------------

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	//----------------------------------------------
	//----------------------------------------------
	//----------------------------------------------

	//-------UnMarshal------------------------------

	s := `[{"Firstname":"James","Lastname":"Bond","Age":99},{"Firstname":"Hamid","Lastname":"Hosseinzadeh","Age":30}]`

	type Person struct {
		Firstname string `json:"Firstname"`
		Lastname  string `json:"Lastname"`
		Age       int    `json:"Age"`
	}

	//var People []Person
	People := []Person{}

	bbss := []byte(s)

	err = json.Unmarshal(bbss, &People)

	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Printf("Person Number %v \n \t Details:%v ", i, v)
	}

}
