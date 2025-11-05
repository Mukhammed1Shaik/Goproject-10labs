package main

import "myproject10/student"

func main() {
	s := student.Student{
		Name:    "Мухаммед",
		Faculty: "Информатика",
		Course:  3,
	}
	s.Info()
}
