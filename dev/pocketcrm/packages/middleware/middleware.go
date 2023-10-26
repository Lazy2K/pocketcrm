package middleware

import (
	"fmt"
	"net/http"
)

// Logging middleware function
func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nRoute: %s\nMethod: %s\n", r.URL.Path, r.Method)
		next(w, r)
	}
}

// User authentication middleware
func Auth(){
	
}