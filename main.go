package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "GET!")
		case http.MethodPost:
			fmt.Fprintf(w, "POST!")
		default:
			fmt.Fprintf(w, "Method not Allowed")
		}
	})

}
