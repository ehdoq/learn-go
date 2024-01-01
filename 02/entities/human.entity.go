package entities

import "fmt"

type Human struct {
	Name   string
	Gender string
}

func NewHuman(name string, gender string) Human {
	human := Human{name, gender}
	return human
}

func (human Human) ToString() string {
	return fmt.Sprintf("name: %s\ngender: %s", human.Name, human.Gender)
}
