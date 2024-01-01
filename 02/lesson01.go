package main

import (
	"fmt"
	"learn-golang/entities"
)

func main() {
	student := entities.Student{
		Id:    "st01",
		Score: 7.8,
		Human: entities.Human{
			Name:   "Name 1",
			Gender: "male",
		},
	}
	fmt.Println("Student Info")
	fmt.Println(student.ToString())

	employee := entities.Employee{
		Id:     "e01",
		Salary: 1000,
		Human: entities.Human{
			Name:   "Name 2",
			Gender: "female",
		},
	}
	fmt.Println("Employee Info")
	fmt.Println(employee.ToString())
}
