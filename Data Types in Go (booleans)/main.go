package main

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {

	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   4000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   6000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {

		if x.Age >= 30 {
			println(x.Name, "is 30 or older!")
		} else {
			println(x.Name, "is under 30!")
		}

		if x.Age >= 30 && x.Salary > 5000 {
			println(x.Name, "makes more than 5000 and is over 30")
		} else {
			println("Either", x.Name, "makes less than 5000 OR is under 30")
		}

		if x.Age >= 30 || x.Salary > 5000 {
			println(x.Name, "makes more than 5000 or is over 30")
		} else {
			println(x.Name, "makes less than 5000 AND is under 30")
		}
	}
}
