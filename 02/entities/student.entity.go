package entities

import "fmt"

type Student struct {
	Id    string
	Score float64
	Human Human
}

func NewStudent(id string, score float64, human Human) Student {
	student := Student{id, score, human}
	return student
}

func (student Student) ToString() string {
	return fmt.Sprintf("id: %s\nscore: %0.2f\n%s", student.Id, student.Score, student.Human.ToString())
}
