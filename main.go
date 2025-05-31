	package main

	import (
		"github.com/gin-gonic/gin"
	 	"github.com/dev-tams/go-auth/profile"
	)

	func main() {

		r := gin.Default()
		r.POST("/register", profile.RegisterUser)
		r.POST("/login", profile.Login)
		r.Run(":8080")
	}

