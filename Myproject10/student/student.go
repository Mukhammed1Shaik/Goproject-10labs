package student

import "fmt"

type Student struct {
	Name    string
	Faculty string
	Course  int
}

func (s Student) Info() {
	fmt.Printf("%s (%s факультеті, %d курс)\n", s.Name, s.Faculty, s.Course)
}
