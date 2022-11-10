package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "GET!")
		case http.MethodPost:
			fmt.Fprintf(w, "POST!")
			w.WriteHeader(http.StatusOK)
		default:
			fmt.Fprintf(w, "Method not Allowed")
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

}
