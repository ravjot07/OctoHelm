package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// helloHandler responds with a greeting message for Service B.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello from Service B at octo!"
	fmt.Fprintln(w, message)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	http.HandleFunc("/hello", helloHandler)
	log.Printf("Service B: Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
