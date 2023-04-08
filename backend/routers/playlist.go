package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"example.com/backend/controllers"
	"example.com/backend/middlewares"
)

func RegisterPlaylistRoutes(router *mux.Router) {
	router.Handle("/api/playlist/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetAllPlaylists))).Methods("GET")
	router.Handle("/api/playlist/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetPlaylist))).Methods("GET")
	router.Handle("/api/playlist/add/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AddPlaylist))).Methods("POST")
	router.Handle("/api/playlist/{userid}/{playlistid}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeletePlaylist))).Methods("DELETE")
	router.Handle("/api/playlist/addsong/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AddPlaylistSong))).Methods("POST")
	router.Handle("/api/playlist/deletesong/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeletePlaylistSong))).Methods("DELETE")
}
