package main

import (
	"fmt"
	"log"
	"github.com/dev-tams/go-auth/auth"
)

func main() {
	//Register
	user, err := auth.RegisterUser("Tammy", "test@mail.com", "securepass")
	if err != nil{
		log.Fatal("Register error", err)
	}
	fmt.Println("Registered User", user)

	//Login
	token, err := auth.Login("test@mail.com", "securepass")
	if err != nil {
		log.Fatal("Login error:", err)
	}
	fmt.Println("JWT token:", token)


	//validate JWT
	userID, err := auth.ValidateJWT(token)
	if err != nil {
		fmt.Println("Token invalid:", err)
	} else {
		fmt.Println("Token valid for user ID:", userID)
	}

}
