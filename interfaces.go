package main

/*
Interfaces allows us to do polymorphism
Interfaces allows us to define behaviour
A value can be of more than one type
If you have this method , you are my type, thats interface
Staffs

· Principal| Teachers | Non-Academic Staffs | Students | Courses | Classes | Applicant
	- A teacher can grade a student in a course he or she teaches.
	- A student can take a course.
	- The principal can expel a student etc.
	- A principal can admit applicants based on age.
 */

type person struct {
	name string
	id float64
	gender [] string
	age float64
	designation [] string
	address string
}

//and i want to write a method the can make the principal
//write in a new student into an array of student struct


type teacher struct {

}

type student struct {

}


func teachCourse()  {

}

func gradeStudents()  {

}
func takeCourse()  {

}

func expelStudent()  {

}

func admitApplicants()  {

}




func main() {

}