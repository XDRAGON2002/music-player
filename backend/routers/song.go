package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"example.com/backend/controllers"
	"example.com/backend/middlewares"
)

func RegisterSongRoutes(router *mux.Router) {
	router.Handle("/api/song/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetAllSongs))).Methods("GET")
	router.Handle("/api/song/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetSong))).Methods("GET")
	router.Handle("/api/song/add/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AddSong))).Methods("POST")
	router.Handle("/api/song/like/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.LikeSong))).Methods("PUT")
}
