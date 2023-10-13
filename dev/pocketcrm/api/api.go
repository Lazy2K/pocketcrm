package api

import (
	"io"
	"fmt"
	"errors"
	"net/http"
)

func getRoot(w http.ResponseWriter, r * http.Request) {
	fmt.Printf("Got /")
	io.WriteString(w, "Hello World!")
}

func StartServer() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":8000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Error")
	}
}
