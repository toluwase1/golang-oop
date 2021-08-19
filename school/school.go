package school

import (
	"fmt"
	applicant "golang-OOP/Applicants"
	"golang-OOP/person"
	"golang-OOP/students"
	"os"
)

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
	Courses
}


type Principal struct {
	person.Person
}

type Courses struct {
	results map[string]int
}


type AdmissionsOffice interface {
	admitApplicant()
	expelStudent()
}



func (p Principal) admitApplicant(applicant applicant.Applicants) students.Student {

	count:=0
	if applicant.Age <10 {
		fmt.Println("Applicant younger than required age")
	} else {
		applicant.
		count++
	}
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
	c.results = map[string]int{"Maths": 90, "English": 45, "Physics": 62, "Chemistry": 29, "Biology": 52}
	resultPage:= "Result breakdown\n"
	for k, v := range c.results {
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