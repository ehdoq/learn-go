package entities

import "fmt"

type Employee struct {
	Id     string
	Salary float64
	Human  Human
}

func NewEmployee(id string, salary float64, human Human) Employee {
	employee := Employee{id, salary, human}
	return employee
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("id: %s\nsalary: %0.2f\n%s", employee.Id, employee.Salary, employee.Human.ToString())
}
