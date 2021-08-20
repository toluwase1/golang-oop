package students

import (
	"golang-OOP/person"
	"golang-OOP/school"
)



type Student struct {
	Offence bool
	person.Person
	//courses.Courses
	Results map[string]int
}

func NewStudent(offence bool, person person.Person, results map[string]int) *Student {
	return &Student{Offence: offence, Person: person, Results: results}
}





func AddToStudentList (applicants school.Applicants) (string, [] Student) {
	 studentList:= []Student {}
	stud:= Student{
		Offence: false,
		Person:  person.Person{},
	}
	if applicants.Age<10 {
		error:= "You are too young"
		return error, nil
	} else {
		studentList=append(studentList, stud)
		return "", studentList
	}
}


func (s Student) takeCourse()  {
	applicant := school.Applicants{}
 	AddToStudentList(applicant)
}



//type Student struct {
//	employees.Employees
//}
//
//func (e Student) abc(a employees.Employees)  {
//
//}
