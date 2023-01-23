package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aymenhta/quitter/config"
)

type CreatePostReq struct {
	Content string `json:"content"`
}

type CreatePostRes struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	PostedAt string `json:"postedAt"`
	UserId   int    `json:"userId"`
}

// What to return after creating a record
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// ! Here make sure to get the id of the current authenticated user
	userId := 1

	// GET The request body
	dto := &CreatePostReq{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dto)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, "Could not decode request's body", http.StatusInternalServerError)
		return
	}

	// TODO: validate the request data

	// INSERT THE RECORD INTO THE DATABASE
	id, postedAt, err := config.G.PostsModel.Insert(dto.Content, userId)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// SEND The response back
	res := CreatePostRes{ID: id, Content: dto.Content, PostedAt: postedAt, UserId: 1}
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(res)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, "Could not marshal json", http.StatusInternalServerError)
		return
	}
}
