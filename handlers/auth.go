package handlers

import (
	"net/http"

	"github.com/aymenhta/quitter/config"
	"github.com/aymenhta/quitter/helpers"
	"github.com/aymenhta/quitter/helpers/validator"
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
	// GET the request body
	dto := &SignUpReq{}
	helpers.DecodeReq(w, r, dto)

	// Validate req body
	v := validator.NewValidator()
	v.Check(dto.Username != "", "username", "must be provided")
	v.Check(len(dto.Username) < 3 || len(dto.Username) > 64, "username", "must be between 3 to 64 bytes long")
	v.Check(dto.Email != "", "email", "must be provided")
	v.Check(validator.Matches(dto.Email, validator.EmailRx), "email", "must be in the correct format")
	v.Check(dto.Password != "", "password", "must be provided")
	v.Check(len(dto.Password) < 6 || len(dto.Password) > 64, "password", "must be atleast 6 bytes long")
	if !v.Valid() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		helpers.EncodeRes(w, v.Errors)
		return
	}

	// Add user to the database
	id, err := config.G.UsersModel.Create(dto.Username, dto.Email, dto.Password)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response
	res := AuthRes{Id: id, Username: dto.Username, Email: dto.Email, Token: "This-Auth-token"}
	w.WriteHeader(http.StatusCreated)
	helpers.EncodeRes(w, res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the sign in"))
}
