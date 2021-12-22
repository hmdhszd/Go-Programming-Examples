package main

import (
	"fmt"
	"log"
)

//------------------------------------------------------
// Basic types: number, strings, booleans

var myInt int

//var myInt16 int16
//var myInt32 int32
//var myInt64 int64

// positive values or zero
var myUint uint

var myFloat float32
var myFloat64 float64

//------------------------------------------------------
// Aggregate types: array , struct

type myCarStruct struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

//------------------------------------------------------

//------------------------------------------------------

//------------------------------------------------------

func main() {

	// Basic types: numbers, strings, booleans

	myInt = 234
	myUint = 10
	myFloat = 10.34
	myFloat64 = 100.234

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "hamid"
	log.Println(myString)

	myString = "Updated myString"
	log.Println(myString)

	var myBool = true
	myBool = false
	log.Println(myBool)

	//------------------------------------------------------
	// Aggregate types: array , struct

	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is :", myStrings[0])

	//------------------------------------------------------

	var myCar myCarStruct

	myCar.NumberOfTires = 4
	myCar.Luxury = false
	myCar.BucketSeats = true
	myCar.Make = "BMW"
	myCar.Model = "740LI"
	myCar.Year = 2021

	//------------------------------------------------------

	myOtherCar := myCarStruct{

		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   false,
		Make:          "BENZ",
		Model:         "S500",
		Year:          2020,
	}

	fmt.Printf("my car is a %d %s %s", myOtherCar.Year, myOtherCar.Make, myOtherCar.Model)

	//------------------------------------------------------
	//------------------------------------------------------
	//------------------------------------------------------
	//------------------------------------------------------

	type person struct {
		firstName string
		lastName  string
		age       int
		friends   []string
	}

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		age:       20,
		friends:   []string{"First Freind of P1", "Second Freind of P1"},
	}

	p2 := person{
		firstName: "Hamid",
		lastName:  "Hosseinzadeh",
		age:       30,
		friends:   []string{"First Freind of P2", "Second Freind of P2"},
	}

	type SercetPerson struct {
		person
		LicenceToKill bool
	}

	sa1 := SercetPerson{

		person: person{
			firstName: "Miss",
			lastName:  "MoneyPenny",
			age:       27,
			friends:   []string{"First Freind of sa1", "Second Freind of sa1"},
		},

		LicenceToKill: true,
	}

	fmt.Println(p1.firstName, p1.lastName, p1.age, p1.friends)
	fmt.Println(p2.firstName, p2.lastName, p2.age, p2.friends)

	fmt.Println(sa1.firstName, sa1.lastName, sa1.age, sa1.friends, sa1.LicenceToKill)
	fmt.Println(sa1.person.firstName, sa1.person.lastName, sa1.person.age, sa1.person.friends, sa1.LicenceToKill)

	//------------------------------------------------------
	//------------------------------------------------------
	//------------------------------------------------------
	//-------------------Anonymous Struct-------------------
	//     when you want to use a STRUCT only in one area

	anonymousPerson := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Nobody's FirstName",
		lastName:  "Nobody's LastName",
		age:       100,
	}

	fmt.Println(anonymousPerson)
}
