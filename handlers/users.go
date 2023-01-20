package handlers

import (
	"encoding/json"
	"net/http"
)

type (
	UserSignupReqDto struct {
		Username, Email, Password string
	}

	UserSigninReqDto struct {
		Email, Password string
	}
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var dto UserSignupReqDto
	err := json.Unmarshal([]byte(""), &dto)
	if err != nil {
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) { w.Write([]byte("This is the sign in")) }
