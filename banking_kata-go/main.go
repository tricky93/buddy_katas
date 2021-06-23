package main

type Account struct {
	balance float64
}

func main() {

}

func createNewAccount() Account {
	a := Account{}

	return a
}

func (a *Account) getBalance() (b int) {
	return
}

func (a *Account) deposit(dA float64) (err string) {
	if dA < 0.01 {
		err = "invalid deposit amount"
		return
	}

	a.balance += dA

	return
}
