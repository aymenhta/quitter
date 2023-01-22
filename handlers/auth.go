package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aymenhta/quitter/config"
)

type (
	SignUpReq struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	SignInReq struct {
		Email, Password string
	}

	AuthRes struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// GET the request body
	dto := &SignUpReq{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dto)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, "Could not decode the request body", http.StatusInternalServerError)
		return
	}
	// TODO: Validate json

	// Add user to the database
	id, err := config.G.UsersModel.Insert(dto.Username, dto.Email, dto.Password)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response
	res := AuthRes{Id: id, Username: dto.Username, Email: dto.Email, Token: "This-Auth-token"}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(res)
	if err != nil {
		http.Error(w, "Could not marshal json", http.StatusInternalServerError)
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This is the sign in"))
}
