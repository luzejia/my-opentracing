package main

import (
	"fmt"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Println("my request id: ", r.Header.Get("X-Request-Id"), ";", "my span id:", r.Header.Get("X-B3-Spanid"))
	fmt.Fprintf(w, "Hello from service D")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
