package wallet

import (
    "testing"
)

func TestWallet(t *testing.T) {
    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10.0))

        assertBalance(t, wallet, Bitcoin(10.0))
    })

    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20.0)}
        err := wallet.Withdraw(Bitcoin(10.0))

        assertBalance(t, wallet, Bitcoin(10.0))
        assertNoError(t, err)
    })

    t.Run("Withdraw insufficient funds", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(10.0)}
        err := wallet.Withdraw(Bitcoin(20.0))

        assertBalance(t, wallet, Bitcoin(10.0))
        assertError(t, err, InsufficientFundsError)
    })
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin)  {
    t.Helper()
    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertError(t *testing.T, got error, want error)  {
    t.Helper()
    if got == nil {
        t.Fatal("Expected an error, got none")
    }

    if got != want {
        t.Errorf("wanted %s, got %s", want, got)
    }
}

func assertNoError(t *testing.T, got error)  {
    t.Helper()
    if got != nil {
        t.Errorf("Unexpected error %s", got)
    }
}


