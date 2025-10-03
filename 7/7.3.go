// Ларионова Арина 363
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type User struct {
	Username string
	Email    string
	Password string
}

func (u *User) SetPassword(password string) {
	h := sha256.Sum256([]byte(password))
	u.Password = hex.EncodeToString(h[:])
}

func (u *User) VerifyPassword(password string) bool {
	h := sha256.Sum256([]byte(password))
	return u.Password == hex.EncodeToString(h[:])
}
func main() {
	userA := User{
		Username: "trtrtr",
		Email:    "trtrtrr@mail.com",
	}
	userA.SetPassword("Baadd00r")
	if userA.VerifyPassword("Baadd00r") {
		fmt.Println("Пароль ввверный")
	} else {
		fmt.Println("Неверный пароль мухахахаха")
	}
}
