package main

import (
	"fmt"
	"github.com/dev-tams/go-auth/auth"
)

func main() {
	// Step 1: Hash a password
	hashed, err := auth.HashPassword("my-secret")
	if err != nil {
		panic(err)
	}
	fmt.Println("Hashed password:", hashed)

	// Step 2: Verify login
	err = auth.CheckPassword(hashed, "my-secret")
	if err != nil {
		fmt.Println("Incorrect password")
	} else {
		fmt.Println("Password is correct")
	}

	token, err := auth.GenerateJWT("12345")
	if err != nil {
		panic(err)
	}
	fmt.Println("JWT token:", token)

	userID, err := auth.ValidateJWT(token)
	if err != nil {
		fmt.Println("Token invalid:", err)
	} else {
		fmt.Println("Token valid for user ID:", userID)
	}

}
