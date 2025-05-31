package profile

import(
	"encoding/json"
	"net/http"
	"strings"
	"github.com/dev-tams/go-auth/auth"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request){

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := auth.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"user_id": userID})
}