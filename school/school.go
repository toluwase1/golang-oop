package school

import (
	"fmt"
	"golang-OOP/person"
)
type Applicants struct {
	id   int
	Age  int
	person.Person
}

func NewApplicants(id int, age int, person person.Person) *Applicants {
	return &Applicants{id: id, Age: age, Person: person}
}



//type Applicants struct {
//	applicant.Applicants
//}

type StaffType struct {
	teacher bool
	nonAcademic bool
}

type Class struct {
	Jss1 string
	Jss2 string
	Jss3 string
}

type Staff struct {
	person.Person
	StaffType
}

type Hundred struct {
	Class
}

//check something here
type Teacher struct {
	Staff
	name string
	Class
}

func NewTeacher(staff Staff, name string, class Class) *Teacher {
	return &Teacher{Staff: staff, name: name, Class: class}
}


type AdmissionsOffice interface {
	admitApplicant()
	expelStudent()
}

type Principal struct {
	person.Person
}

func NewPrincipal(person person.Person) *Principal {
	return &Principal{Person: person}
}


//func applicants(name string) Applicants {
//	newApplicant:= Applicants{
//		Age:    0,
//		Person: person.Person{},
//	}
//	return newApplicant
//}

func (p Principal) AdmitApplicant(applicant *Applicants) Student {
	stud:= NewStudent(false, applicant.Person, map[string]int{})
	count:=0
	//name:=""
	//applicant1:=NewApplicants()
	if applicant.Age <10 {
		fmt.Println("Applicant younger than required age")
	} else {
		count++
	}
	return stud
}

func (p Principal) AddToStudentList (applicants *Applicants, student *[]Student) (string, []Student) {
	//studentList:= []Student{}
	//stud:= Student{
	//	Offence: false,
	//	Person:  person.Person{},
	//}
	stud:= NewStudent(false, applicants.Person, map[string]int{})
	if applicants.Age<10 {
		error:= "You are too young"
		return error, nil
	} else {
		*student=append(*student, stud)
		fmt.Println(student)
		return "", *student
	}
}

/*
teacher takes courses
a method that takes in teacher name and produces the shows the courses he/she takes
takes in courses and produces names and classes of teachers

student takes courses
a method that takes in teacher name and produces the shows the courses he/she takes
takes in courses and produces names and classes of teachers
*/

//func (t Teacher) teachCourse(c courses.Courses)  {
//	for i,j:= range c.Class {
//
//	}
//}


//func (s *bill) save()  {
//	data:= [] byte (b.format()) //format and put into a byteslice, then save into data
//	err:= os.WriteFile("bills/"+b.name+".txt", data,0666)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Print("Bill was saved to the file")
//}



func (p Principal) expelStudent(student *Student)  {
	if student.Offence ==true {
		fmt.Println("You have been expelled from this school")
	}
}

func (t Teacher) GradeStudent(c Student) string {
//SUPPLY TEACHER INFO AND STUDENT INFO BEFORE GRADING
	resultPage:= "Result Page \n"
	resultPage+=fmt.Sprintf("Class Teacher %s \n", t.name)
	resultPage+=fmt.Sprintf("Student Name %s \n", c.Name)
	resultPage+=fmt.Sprintf("Student Name %s \n", c.Name)
	for k, v := range c.Results {
		v+=v
		if v >= 70 {
			resultPage+=fmt.Sprintf("You scored %i+A in %s \n", v, k)
		} else if v > 59 && v < 70 {
			resultPage+=fmt.Sprintf("You scored %i+B in %s \n", v, k)
		} else if v > 49 && v < 60 {
			resultPage+=fmt.Sprintf("You scored %i+C in %s \n", v, k)
		} else if v > 39 && v < 49 {
			resultPage+=fmt.Sprintf("You scored %i+D in %s \n", v, k)
		} else {
			resultPage+=fmt.Sprintf("You failed with %i+F in %s \n", k, v)
		}
		aggregate:=v/5
		if aggregate>49 {
			resultPage+=fmt.Sprintf("You have passed the semester with an aggregate of %i \n", aggregate)
		} else {
			resultPage+=fmt.Sprintf("You have failed the semester with an aggregate of %i \n", aggregate)
		}
	}
	resultPage+=fmt.Sprintf("%-25v ...$%0.2f")
	return resultPage
}
