package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	msg := "Hello, World!\n"
	port := 8080

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, msg)
	})
	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
