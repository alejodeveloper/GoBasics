package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	psswd := `password123`
	encryptedPsswd, err := bcrypt.GenerateFromPassword([]byte(psswd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encryptedPsswd)

	loginPsswd := `password123`
	err = bcrypt.CompareHashAndPassword(encryptedPsswd, []byte(loginPsswd))
	if err != nil {
		fmt.Println(err)
	}
}
