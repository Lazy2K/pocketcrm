package dashboard

import (
	"embed"
	"fmt"
	"net/http"
	"strings"
)

var distDir = "ui/pocketcrm-ui/dist"

// Reformat all this is into neeter stuff, maybe a middleware module
func rootPath(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = fmt.Sprintf("/%s/", distDir)
		} else {
			b := strings.Split(r.URL.Path, "/")[0]
			if b != distDir {
				r.URL.Path = fmt.Sprintf("/%s%s", distDir, r.URL.Path)
			}
		}
		h.ServeHTTP(w,r)
	})
}

func StartServer(port string, files embed.FS) {
	// Setup server mux
	mux := http.NewServeMux()

	// I don't yet understand why this works but it does....
	fileHandler := http.FileServer(http.FS(files))
	rootFileHandler := rootPath(fileHandler)
	mux.Handle("/", rootFileHandler)

	// Listen on port
	http.ListenAndServe(port, mux)
}