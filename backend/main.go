package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"example.com/backend/routers"
)

func main() {
	router := mux.NewRouter()
	routers.RegisterSongRoutes(router)
	routers.RegisterUserRoutes(router)
	routers.RegisterPlaylistRoutes(router)

   credentials := handlers.AllowCredentials()
   methods := handlers.AllowedMethods([]string{"POST", "GET", "DELETE"})
   origins := handlers.AllowedOrigins([]string{"www.example.com", "http://localhost:3000"})
   fmt.Println("Listening on port 5000...")
   log.Fatal(http.ListenAndServe(":5000", handlers.CORS(credentials, methods, origins)(router)))


}
