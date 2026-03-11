package main

import "fmt"

type Bankaccount struct {
	accountholder string
	balance       float64
}

func (b *Bankaccount) Deposit(amt float64) {

	b.balance = b.balance + amt
	fmt.Println("Balance after deposting", b.balance, "with deposit amt", amt)

}

func (b *Bankaccount) Withdraw(amt float64) {

	if b.balance < amt {

		fmt.Println("Balance is insufficient")
	} else {
		b.balance = b.balance - amt
		fmt.Println("Balance after withdrawl", b.balance, "with amt", amt)
	}

}

func main() {
	bnkacc := Bankaccount{
		accountholder: "Rakesh",
		balance:       123.3,
	}
	bnkacc.Deposit(100)
	bnkacc.Withdraw(500)

	fmt.Println("Program finished")

}
