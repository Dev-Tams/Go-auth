	package main

	import (
		"fmt"
		"log"
		"net/http"
		"go-auth/profile"
	)

	func main() {

		http.HandleFunc("/register", RegisterUser)
		http.HandleFunc("/login", LoginUser)
		http.HandleFunc("/profile", ProfileHandler)

		fmt.Println("✅ Server running at http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

