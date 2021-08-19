package students

import (
	"golang-OOP/person"
	"golang-OOP/school"
)



type Student struct {
	Offence bool
	grades  [] string
	person.Person
	school.Courses
}

func student() Student {

	stud:= Student{
		Offence: false,
		grades:  nil,
		Person:  person.Person{},
		Courses: school.Courses{},
	}
	studentList:= [] [] Student {}
	return stud
}


func (s Student) takeCourse()  {

}



//type Student struct {
//	employees.Employees
//}
//
//func (e Student) abc(a employees.Employees)  {
//
//}
