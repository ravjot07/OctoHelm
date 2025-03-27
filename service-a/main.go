package main 

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// helloHandler responds with a greeting message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello from Service A from octo!"
	fmt.Fprint(w, message)
}

func main(){
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/hello", helloHandler)
	log.Printf("Service A: Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}