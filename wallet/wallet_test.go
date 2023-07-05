package wallet

import (
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		testhelpers.AssertCorrectString(t, got.String(), want.String())
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		testhelpers.AssertNoError(t, err)
		testhelpers.AssertCorrectString(t, got.String(), want.String())
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(0)

		testhelpers.AssertError(t, err, ErrInsufficientFunds.Error())
		testhelpers.AssertCorrectString(t, got.String(), want.String())
	})

}
