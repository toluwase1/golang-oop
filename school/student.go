package school

import (
	"golang-OOP/person"
)



type Student struct {
	Offence bool
	person.Person
	Teacher
	Results map[string]int
	Courses
}

func NewStudent2(offence bool, person person.Person, teacher Teacher, results map[string]int, courses Courses) *Student {
	return &Student{Offence: offence, Person: person, Teacher: teacher, Results: results, Courses: courses}
}


func NewStudent(offence bool, person person.Person, results map[string]int) Student {
	return Student{Offence: offence, Person: person, Results: results}
}





//func (s Student) takeCourse()  {
//	applicant := school.Applicants{}
// 	AddToStudentList(applicant)
//}



//type Student struct {
//	employees.Employees
//}
//
//func (e Student) abc(a employees.Employees)  {
//
//}
