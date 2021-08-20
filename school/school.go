package school

import (
	"fmt"
	"golang-OOP/person"
	"golang-OOP/students"
)
type Applicants struct {
	name string
	id   int
	Age  int
	person.Person
}

func NewApplicants(name string, id int, age int, person person.Person) *Applicants {
	return &Applicants{name: name, id: id, Age: age, Person: person}
}

//type Applicants struct {
//	applicant.Applicants
//}

type StaffType struct {
	academic bool
	nonAcademic bool
}

type Class struct {
	hundredLevel string
	twoHundredLevel string
	threeHundredLevel string
}

type Staff struct {
	person.Person
	StaffType
}

type Hundred struct {
	Class
}

//check something here
type AcademicStaff struct {
	Staff
}





type AdmissionsOffice interface {
	admitApplicant()
	expelStudent()
}

type Principal struct {
	person.Person
}


func applicants(name string) Applicants {
	newApplicant:= Applicants{
		Age:    0,
		Person: person.Person{},
	}
	return newApplicant
}

func (p Principal) AdmitApplicant(applicant Applicants) *students.Student {
	stud:= students.NewStudent(false, applicant.Person, map[string]int{})
	count:=0
	name:=""
	applicant1:=applicants(name)
	if applicant.Age <10 {
		fmt.Println("Applicant younger than required age")
	} else {
		students.AddToStudentList(applicant1)
		count++
	}
	return stud
}




//func (s *bill) save()  {
//	data:= [] byte (b.format()) //format and put into a byteslice, then save into data
//	err:= os.WriteFile("bills/"+b.name+".txt", data,0666)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Print("Bill was saved to the file")
//}



func (p Principal) expelStudent(student *students.Student)  {
	if student.Offence ==true {
		fmt.Println("You have been expelled from this school")
	}
}

func (t AcademicStaff) gradeStudent(c students.Student) string {
	//composition
	resultPage:= "Result breakdown for \n"
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

func (t AcademicStaff) teachCourse()  {
	fmt.Println("I am teaching a course")
}