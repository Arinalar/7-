// Ларионова Арина 363
package main

import (
    "errors"
    "fmt"
)

// BankAccount представляет структуру банковского счёта.
type BankAccount struct {
    accountNumber string  // Номер счёта
    holderName    string  // Имя держателя счёта
    balance       float64 // Баланс счёта
}

// Deposit пополняет счёт на указанную сумму.
func (ba *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("число для пополнения должно быть положительным")
    }
    ba.balance += amount
    return nil
}

// Withdraw списывает указанную сумму со счёта. Возвращает ошибку, если средств недостаточно.
func (ba *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("число для списания должно быть положительным")
    }
    if amount > ba.balance {
        return errors.New("недостаточно средств для списания")
    }
    ba.balance -= amount
    return nil
}

// GetBalance возвращает текущий баланс счёта.
func (ba *BankAccount) GetBalance() float64 {
    return ba.balance
}

// Функция main - точка входа в программу.
func main() {
    acc := BankAccount{
        accountNumber: "1",
        holderName:    "Арина",
        balance:       0,
    }

    // Пополнение счёта
    if err := acc.Deposit(100); err != nil {
        fmt.Println("Ошибка:", err)
    }

    // Списание со счёта
    if err := acc.Withdraw(50); err != nil {
        fmt.Println("Ошибка:", err)
    }

    // Печать текущего баланса
    fmt.Printf("Текущий баланс: %.2f\n", acc.GetBalance())
}
