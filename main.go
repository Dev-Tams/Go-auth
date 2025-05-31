	package main

	import (
		"fmt"
		"log"
		"net/http"
	 	"github.com/dev-tams/go-auth/profile"
	)

	func main() {

		http.HandleFunc("/register", profile.RegisterUser)
		http.HandleFunc("/login", profile.LoginUser)
		http.HandleFunc("/profile", profile.ProfileHandler)

		fmt.Println("✅ Server running at http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

