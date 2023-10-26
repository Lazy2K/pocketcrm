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

func login(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Login Page\n")
}

func account(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Account Page\n")
}

func StartServer(port string) {
	// Setup server mux
	mux := http.NewServeMux()

	// Un-authenticated rotues
	mux.HandleFunc("/login", middleware.Log(login))

	// Authenticated routes
	mux.HandleFunc("/", middleware.Auth(root))
	mux.HandleFunc("/account", middleware.Auth(account))

	// Listen on port
	http.ListenAndServe(port, mux)
}
