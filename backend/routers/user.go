package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"example.com/backend/controllers"
	"example.com/backend/middlewares"
)

func RegisterUserRoutes(router *mux.Router) {
	router.Handle("/api/user/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetUsers))).Methods("GET")
	router.Handle("/api/user/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetUser))).Methods("GET")
	router.Handle("/api/user/add/", http.HandlerFunc(controllers.AddUser)).Methods("POST")
	router.Handle("/api/user/login/", http.HandlerFunc(controllers.LoginUser)).Methods("POST")
}
