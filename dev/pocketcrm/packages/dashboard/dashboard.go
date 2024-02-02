package dashboard

import (
	"net/http"
)




func StartServer(port string) {
	// Setup server mux
	mux := http.NewServeMux()

	// I don't yet understand why this works but it does....
	fileHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/pocketcrm-ui/dist")))
	mux.Handle("/static/", fileHandler)

	// Listen on port
	http.ListenAndServe(port, mux)
}