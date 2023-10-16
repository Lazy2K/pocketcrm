package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"pocketcrm/database"
)

func getRoot(w http.ResponseWriter, r * http.Request) {
	database.Query()
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
