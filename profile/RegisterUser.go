package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dev-tams/go-auth/auth"
)

func RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	
	if len(req.Email) <= 5 {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid mail"})
	return
	}
	
	if len(req.Password) <= 7 {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Password too short"})
	return
}


	// Register user
	user, err := auth.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Success
	c.JSON(http.StatusOK, user)
}
