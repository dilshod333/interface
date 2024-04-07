package main

import (
	"fmt"
)

type Student struct {
	name   string
	id     int
	grades []float64
}

func NewStudent(name string, id int) *Student {
	return &Student{
		name:   name,
		id:     id,
		grades: []float64{},
	}
}

func (s *Student) AddGrade(grade float64) {
	s.grades = append(s.grades, grade)
}

func (s *Student) CalculateAverage() float64 {
	sum := 0.0
	for _, grade := range s.grades {
		sum += grade
	}
	if len(s.grades) == 0 {
		return 0.0
	}
	return sum / float64(len(s.grades))
}

func (s *Student) DisplayInfo() {
	fmt.Printf("Student Name: %s\n", s.name)
	fmt.Printf("Student ID: %d\n", s.id)
	fmt.Printf("Grades: %v\n", s.grades)
	fmt.Printf("Average Grade: %.2f\n", s.CalculateAverage())
}

func main() {

	student := NewStudent("Dilshod", 12345)

	student.AddGrade(85.5)
	student.AddGrade(90.0)
	student.AddGrade(78.5)

	student.DisplayInfo()
}
