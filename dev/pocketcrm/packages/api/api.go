package api

import (
	"io"
	"net/http"
	"pocketcrm/packages/database"
	"pocketcrm/packages/middleware"
)

func root(w http.ResponseWriter, r *http.Request) {
	database.Query()
	io.WriteString(w, "Hello World!\n")
}

func account(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Account Page\n")
}

func StartServer(port string) {
	// Setup server mux
	mux := http.NewServeMux()

	// Mux routes
	mux.HandleFunc("/", middleware.Log(root))
	mux.HandleFunc("/account", middleware.Log(account))

	// Listen on port
	http.ListenAndServe(port, mux)
}
