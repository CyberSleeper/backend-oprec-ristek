package routes

import (
	"github.com/CyberSleeper/backend-oprec-ristek/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/{postId}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/{postId}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/{postId}", controllers.DeletePost).Methods("DELETE")
}
