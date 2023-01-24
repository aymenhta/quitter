package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aymenhta/quitter/config"
	"github.com/aymenhta/quitter/helpers"
	"github.com/aymenhta/quitter/models"
	"github.com/julienschmidt/httprouter"
)

type (
	CreatePostReq struct {
		Content string `json:"content"`
	}

	CreatePostRes struct {
		ID       int    `json:"id"`
		Content  string `json:"content"`
		PostedAt string `json:"postedAt"`
		UserId   int    `json:"userId"`
	}

	PostsListRes struct {
		ID       int    `json:"id"`
		Content  string `json:"content"`
		PostedAt string `json:"postedAt"`
		UserId   int    `json:"userId"`
		Username string `json:"username"`
	}

	PostDetailsRes struct {
		ID       int    `json:"id"`
		Content  string `json:"content"`
		PostedAt string `json:"postedAt"`
		UserId   int    `json:"userId"`
		Username string `json:"username"`
	}
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := config.G.PostsModel.GetPosts()
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back the response
	var res []PostsListRes
	for _, v := range posts {
		res = append(res, PostsListRes{
			ID:       v.ID,
			Content:  v.Content,
			PostedAt: helpers.ToHumanTimeStamp(v.PostedAt),
			UserId:   v.UserId,
			Username: v.Username,
		})
	}
	w.WriteHeader(200)
	helpers.EncodeRes(w, res)
}

func PostDetails(w http.ResponseWriter, r *http.Request) {
	// GET :id value from the url
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// GET Post Details From the database
	postDetails, err := config.G.PostsModel.GetPost(id)
	if err != nil {
		config.G.ErrorLog.Println(err)
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Could not get post details", http.StatusInternalServerError)
		}
		return
	}

	// Send back the response
	res := PostDetailsRes{
		ID:       postDetails.ID,
		Content:  postDetails.Content,
		PostedAt: helpers.ToHumanTimeStamp(postDetails.PostedAt),
		UserId:   postDetails.UserId,
		Username: postDetails.Username,
	}
	w.WriteHeader(200)
	helpers.EncodeRes(w, res)
}

// What to return after creating a record
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// ! Here make sure to get the id of the current authenticated user
	userId := 1

	// GET The request body
	dto := &CreatePostReq{}
	helpers.DecodeReq(w, r, dto)

	// TODO: validate the request data

	// INSERT THE RECORD INTO THE DATABASE
	id, postedAt, err := config.G.PostsModel.Create(dto.Content, userId)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// SEND The response back
	res := CreatePostRes{
		ID:       id,
		Content:  dto.Content,
		PostedAt: helpers.ToHumanTimeStamp(postedAt),
		UserId:   1,
	}
	w.WriteHeader(201)
	helpers.EncodeRes(w, res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// GET :id value from the url
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// DELETE Post
	err = config.G.PostsModel.DeletePost(id)
	if err != nil {
		config.G.ErrorLog.Println(err)
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Could not get post details", http.StatusInternalServerError)
		}
		return
	}

	msg := fmt.Sprintf("Post %v was deleted successfully", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
