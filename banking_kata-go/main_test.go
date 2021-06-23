package main

// Bank Kata
// Develop an application that can deposit funds, withdraw funds and print a statement of a customer’s
// bank account activity.
// Here's the specification for an acceptance test that expresses the desired behaviour for this:
// Given a client makes a deposit of 1000 on 10/01/2021
// And a deposit of 2000 on 13/01/2021
// And a withdrawal of 500 on 14/01/2021
// When they print their bank statement
// Then they would see:
// Date || Amount || Balance
// 14/01/2012 || -500 || 2500
// 13/01/2012 || 2000 || 3000
// 10/01/2012 || 1000 || 1000

import (
	"errors"
	"fmt"
	"testing"
)

func TestAccountBalance(t *testing.T) {
	a := createNewAccount()
	want := 0.00
	got := a.balance

	if want != got {
		t.Errorf("got %v. want %v", got, want)
	}
}

func TestDeposit(t *testing.T) {
	testCases := []struct {
		desc            string
		depositAmount   float64
		expectedBalance float64
		expectedError   error
	}{{
		desc:            "Success - Deposits 500.00",
		depositAmount:   500.00,
		expectedBalance: 500.00,
	}, {
		desc:            "Success - Deposits 100.00",
		depositAmount:   100.00,
		expectedBalance: 100.00,
	},
		{
			desc:            "Fail - Deposits 100.00",
			depositAmount:   -100.00,
			expectedBalance: 0,
			expectedError:   errors.New("invalid deposit amount"),
		},
	}

	for i := 0; i < len(testCases); i++ {
		a := createNewAccount()
		err := a.deposit(testCases[i].depositAmount)

		if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", testCases[i].expectedError) {
			t.Errorf("expected error %v. actual error %v", testCases[i].expectedError, err)
		}

		if testCases[i].expectedBalance != a.balance {
			t.Errorf("expected balance %v. actual balance %v", testCases[i].expectedBalance, a.balance)
		}
	}

	// a := createNewAccount()
	// a.deposit(500.00)
	// want := 500.00
	// got := a.balance

	// if want != got {
	// 	t.Errorf("got %v. want %v", got, want)
	// }
}
