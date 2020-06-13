package pointers

import (
	"fmt"
	"testing"
)

type Bitcoin int

func (b *Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return fmt.Errorf("insufficient balance: %d", w.balance)
	}
	w.balance -= amt
	return nil
}

func TestWallet(t *testing.T) {

	checkBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got.String(), want.String())
		}

	}

	t.Run("deposit bitcoins", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		checkBalance(t, wallet, 10)
	})

	t.Run("withdraw bitcoins", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		checkBalance(t, wallet, 10)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(20)
		checkBalance(t, wallet, 10)
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
