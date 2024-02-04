package main

import "fmt"

type Student struct {
	ID int
	Name string
	GPA float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
}

func main() {
	student := Student{1, "Enjang Setiawan", 3.33}
	fmt.Println(student.Name)

	//graduate(&student)
	student.graduate() // change from function to method

	fmt.Println(student.Name)
}
