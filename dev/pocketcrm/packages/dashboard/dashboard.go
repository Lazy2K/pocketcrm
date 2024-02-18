package dashboard

import (
	"embed"
	"net/http"
)



func StartServer(port string, files embed.FS) {
	// Setup server mux
	mux := http.NewServeMux()

	// I don't yet understand why this works but it does....
	fileHandler := http.FileServer(http.FS(files))
	mux.Handle("/", fileHandler)

	// Listen on port
	http.ListenAndServe(port, mux)
}