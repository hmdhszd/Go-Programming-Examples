package staff

import "fmt"

var OverPaidLimit = 75000

var UnderPaidLimit = 40000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStafs []Employee
}

func (e *Office) All() []Employee {
	return e.AllStafs
}

func (e *Office) OverPaid() []Employee {

	var overpaid []Employee

	for _, x := range e.AllStafs {

		if x.Salary > OverPaidLimit {
			fmt.Println("OverPaidLimit is :", OverPaidLimit)
			overpaid = append(overpaid, x)
		}

	}

	return overpaid
}

func (e *Office) UnerPaid() []Employee {

	var underpaid []Employee

	for _, x := range e.AllStafs {

		if x.Salary < UnderPaidLimit {
			fmt.Println("UnderPaidLimit is :", UnderPaidLimit)
			underpaid = append(underpaid, x)
		}

	}

	return underpaid
}
