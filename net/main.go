// https://golang.org/pkg/
package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{message: 'hola'}")
}

func main() {
	http.HandleFunc("/hola", handler)
	http.ListenAndServe(":8000", nil)
}
