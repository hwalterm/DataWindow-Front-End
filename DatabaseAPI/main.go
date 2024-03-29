package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// Hello world, the web server
	log.Println("Listening for requests at http://localhost:8080/hello")

	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe("localhost:8080", nil)
	log.Println("Listening for requests at http://localhost:8080/hello")
	//log.Fatal(http.ListenAndServeTLS("localhost:8080", "localhost.crt", "localhost.key", nil))
}
