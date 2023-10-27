package middleware

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Logging middleware function
func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// End middleware, call next handler
		next(w, r)
	}
}

// User authentication middleware
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nMethod: %s\nRoute: %s\n", r.Method, r.URL.Path)
		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			r.Body.Close()
			if err != nil {
				fmt.Printf("Post Request Body Parse Error\n")
				log.Fatalln(err)
			}
			fmt.Printf("Body: \n%s\n", body)
		}

		// End middleware, call next handler
		next(w, r)
	}
}