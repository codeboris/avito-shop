package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Printf("Listening on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
