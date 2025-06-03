package main

import (
	"log"

	"github.com/dev-tams/go-auth/profile"
	"github.com/dev-tams/go-auth/auth"
	"github.com/gin-gonic/gin"
)

func main() {

	//init db
	if err := auth.InitDB(); err != nil {
        log.Fatal("Failed to connect to DB: ", err)
    }

	r := gin.Default()
	r.POST("/register", profile.RegisterUser)
	r.POST("/login", profile.Login)
	r.Run(":8080")
}
