package main

import (
	"fmt"
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 60000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 65000, FullTime: true},
	{FirstName: "Alan", LastName: "Andreson", Salary: 50000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {

	myStaff := staff.Office{
		AllStafs: employees,
	}

	log.Println(myStaff.All())
	fmt.Println()

	log.Println(myStaff.OverPaid())
	fmt.Println()

	staff.OverPaidLimit = 60000

	log.Println(myStaff.OverPaid())
	fmt.Println()

	log.Println(myStaff.UnerPaid())

}
