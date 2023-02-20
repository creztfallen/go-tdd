package poiters

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		erro := wallet.Withdraw(Bitcoin(10))
		confirmBalance(t, wallet, Bitcoin(10))
		confirmNonExistentError(t, erro)
	})
	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		erro := wallet.Withdraw(Bitcoin(100))

		confirmBalance(t, wallet, initialBalance)
		confirmError(t, erro, InsufficientFundsError.Error())
	})
}

func confirmBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()

	if result != expected {
		t.Errorf("resultado %s, expected %s", result, expected)
	}
}

func confirmError(t *testing.T, value error, expected string) {
	t.Helper()
	if value == nil {
		t.Error("esperava um erro mas nenhum ocorreu")
	}

	result := value.Error()

	if result != expected {
		t.Errorf("resultado %s, esperado %s", result, expected)
	}
}

func confirmNonExistentError(t *testing.T, result error) {
	t.Helper()
	if result != nil {
		t.Fatal("erro inesperado recebido")
	}
}
