package main

type BankAccount struct {
	Owner string
	balance float64
}

func (ba BankAccount) Balance() float64 {
	return ba.balance
}

func (ba *BankAccount) Deposit(amount float64)  {
	ba.balance+=amount
}
func (ba *BankAccount) Withdraw()  {

}

//func (a Principal) abc()  {
//
//}


func main() {
	
}