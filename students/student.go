package students

import (
	"golang-OOP/courses"
	"golang-OOP/person"
	"golang-OOP/school"
)



type Student struct {
	Offence bool
	person.Person
	school.Teacher
	Results map[string]int
	courses.Courses
}

func NewStudent2(offence bool, person person.Person, teacher school.Teacher, results map[string]int, courses courses.Courses) *Student {
	return &Student{Offence: offence, Person: person, Teacher: teacher, Results: results, Courses: courses}
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
