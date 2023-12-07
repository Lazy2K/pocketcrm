package web

import (
	"html/template"
	"net/http"
	"path/filepath"
	"pocketcrm/packages/database"
	"pocketcrm/packages/middleware"
)

func root(w http.ResponseWriter, r *http.Request) {
	database.Query()
	base := filepath.Join("static/templates", "base.html")
	temp, _ := template.ParseFiles(base)
	temp.ExecuteTemplate(w, "base", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	// Do nil rn
}


func StartServer(port string) {
	// Setup server mux
	mux := http.NewServeMux()

	// Un-authenticated rotues
	mux.HandleFunc("/login", middleware.Log(login))

	// Authenticated routes
	mux.HandleFunc("/", middleware.Auth(root))

	// Listen on port
	http.ListenAndServe(port, mux)
}