package main

import (
		"fmt"

		"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		return err == nil
}

func main() {
		password := "this-is-the-password-of-a-user-#2468"
		hash, _ := HashPassword(password)

		fmt.Println("Entered Password:", password)
		fmt.Println("Generated Hash:    ", hash)

		match := CheckPasswordHash(password, hash)
		fmt.Println("Does the hash match the password?:   ", match)
}
