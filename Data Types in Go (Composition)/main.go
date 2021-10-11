package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers:", v.NumberOfPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("IsElectric:", c.IsElectric)
	fmt.Println("IsHybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {

	suv := Vehicle{
		NumberOfPassengers: 6,
		NumberOfWheels:     4,
	}

	volvoxc90 := Car{
		Make:       "Volvo",
		Model:      "XC90",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvoxc90.show()

	fmt.Println()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	teslaModelX.show()

	teslaModelX.Vehicle.NumberOfPassengers = 7

	fmt.Println()

	teslaModelX.show()

}
