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
	fmt.Println("Listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
