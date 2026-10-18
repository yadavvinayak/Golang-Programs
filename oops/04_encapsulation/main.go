package main

import "fmt"

type BankAccount struct{ balance int }

func (a *BankAccount) Deposit(amount int) {
    if amount > 0 { a.balance += amount }
}
func (a BankAccount) Balance() int { return a.balance }

func main() {
    account := BankAccount{}
    account.Deposit(250)
    fmt.Println("balance:", account.Balance())
}
