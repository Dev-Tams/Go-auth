		package profile

	import(
		"encoding/json"
		"net/http"
		"github.com/dev-tams/go-auth/auth"
	)

		func LoginUser(w http.ResponseWriter, r *http.Request){
				//check method
			if r.Method != http.MethodPost{
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			var req struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
				http.Error(w, "Invalid Body param", http.StatusBadRequest)
				return
			}

			token, err := auth.Login(req.Email, req.Password)
			if err != nil {
				http.Error(w, "Login failed", http.StatusUnauthorized)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{"token": token})
		}
