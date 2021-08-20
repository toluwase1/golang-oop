package main

import (
	"fmt"
	"golang-OOP/person"
	"golang-OOP/school"
)

//type BankAccount struct {
//	Owner string
//	balance float64
//}
//
//func (ba BankAccount) Balance() float64 {
//	return ba.balance
//}
//
//func (ba *BankAccount) Deposit(amount float64)  {
//	ba.balance+=amount
//}
//func (ba *BankAccount) Withdraw()  {
//
//}
//
////func (a Principal) abc()  {
////
////}


func main() {
	student:=[]school.Student{}
	applicant1:= school.NewApplicants(1, 11, person.Person{Name: "Tommy"})
	principal1:= school.NewPrincipal(person.Person{Name: "Chika", Gender: "Male", Title: "Mr"})
	principal1.AddToStudentList(applicant1, &student)
	fmt.Println("Applicant with information: ", applicant1, "\n", "Has become student with information: ", student, "\n")


}

