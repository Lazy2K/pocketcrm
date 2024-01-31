package api

import (
	"io"
	"net/http"
	"pocketcrm/packages/database"
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
	mux.HandleFunc("/login", login)

	// Authenticated routes
	mux.HandleFunc("/", root)
	mux.HandleFunc("/account", account)

	// I don't yet understand why this works but it does....
	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("ui/pocketcrm-ui/dist")))
	mux.Handle("/static/", fileHandler)

	// Listen on port
	http.ListenAndServe(port, mux)
}
