package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CyberSleeper/backend-oprec-ristek/pkg/models"
	"github.com/CyberSleeper/backend-oprec-ristek/pkg/utils"
	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	newPosts := models.GetAllPosts()
	res, _ := json.Marshal(newPosts)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	Id, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	postDetails, _ := models.GetPostById(Id)
	res, _ := json.Marshal(postDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	createPost := &models.Post{}
	utils.ParseBody(r, createPost)
	b := createPost.CreatePost()
	res, _ := json.Marshal(&b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	Id, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	post := models.DeletePost(Id)
	res, _ := json.Marshal(post)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	updatePost := &models.Post{}
	utils.ParseBody(r, updatePost)
	vars := mux.Vars(r)
	postId := vars["postId"]
	Id, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	postDetails, db := models.GetPostById(Id)
	if updatePost.Caption != "" {
		postDetails.Caption = updatePost.Caption
	}
	db.Save(&postDetails)
	res, _ := json.Marshal(postDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
