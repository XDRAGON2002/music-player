package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/backend/routers"
)

func main() {
	router := mux.NewRouter()
	routers.RegisterSongRoutes(router)
	routers.RegisterUserRoutes(router)
	routers.RegisterPlaylistRoutes(router)
	fmt.Println("Listening on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
