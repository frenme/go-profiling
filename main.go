package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", homeHandler)
	
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Profiling Demo Server\n")
}