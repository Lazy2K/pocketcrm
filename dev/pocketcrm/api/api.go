package api

import (
	"fmt"
	"io"
	"net/http"
	"pocketcrm/api/middleware"
	"pocketcrm/database"
)

func getRoot(w http.ResponseWriter, r * http.Request) {
	database.Query()
	fmt.Printf("Got /\n")
	io.WriteString(w, "Hello World!\n")
	middleware.Log()
}

func StartServer(port string) {
	// Setup server mux
	mux := http.NewServeMux()

	// Mux routes
	mux.HandleFunc("/", getRoot)

	// Listen on port
	http.ListenAndServe(port, mux)
}
