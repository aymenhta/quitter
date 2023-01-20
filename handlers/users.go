package handlers

import (
	"encoding/json"
	"net/http"
)

type UserSignupDto struct {
	username, email, password string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var dto UserSignupDto
	err := json.Unmarshal([]byte(""), &dto)
	if err != nil {
		return
	}
}
