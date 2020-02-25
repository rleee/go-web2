package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello there</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)

	fmt.Println("http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
