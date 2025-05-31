package profile

import(
	"encoding/json"
	"net/http"
	"github.com/dev-tams/go-auth/auth"
	"github.com/gin-gonic/gin"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, c *gin.Context){
	//check method
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, "Invalid Body param", http.StatusBadRequest)
		return
	}

	user, err := auth.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(user)
}