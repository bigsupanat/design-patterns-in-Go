package main

import "fmt"

const (
	Developer = iota
	Manager
	Ceo
)

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFacotry struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFacotry) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFacotry {
	return &EmployeeFacotry{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func NewEmployeeFactory3(role int) func(name string, annualIncome int) *Employee {
	switch role {
	case Developer:
		return func(name string, annualIncome int) *Employee {
			return &Employee{name, "Developer", annualIncome}
		}
	case Manager:
		return func(name string, annualIncome int) *Employee {
			return &Employee{name, "Manager", annualIncome}
		}
	case Ceo:
		return func(name string, annualIncome int) *Employee {
			return &Employee{name, "CEO", annualIncome}
		}
	default:
		panic("unsupported role")
	}
}

func main() {
	developerFactory := NewEmployeeFactory3(Developer)
	managerFactory := NewEmployeeFactory3(Developer)
	ceoFactory := NewEmployeeFactory3(Developer)

	developer := developerFactory("Adam", 60000)
	manager := managerFactory("Jane", 80000)
	ceo := ceoFactory("Sam", 100000)

	fmt.Println(developer)
	fmt.Println(manager)
	fmt.Println(ceo)

	// bossFactory := NewEmployeeFactory2("CEO", 100000)
	// bossFactory.AnnualIncome = 110000
	// boss := bossFactory.Create("Sam")
	// fmt.Println(boss)
}
