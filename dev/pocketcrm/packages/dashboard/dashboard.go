package dashboard

import (
	"embed"
	"net/http"
	"pocketcrm/packages/middleware"
)

// Move this somewhere else i guess
var distDir = "ui/pocketcrm-ui/dist"

func StartServer(port string, files embed.FS) {
	// Setup server mux
	mux := http.NewServeMux()

	// I don't yet understand why this works but it does....
	fileHandler := http.FileServer(http.FS(files))
	rootFileHandler := middleware.RootPath(fileHandler, distDir)
	mux.Handle("/", rootFileHandler)

	// Listen on port
	http.ListenAndServe(port, mux)
}