package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if len(os.Args) == 2{
		port = os.Args[1]
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	// Serves at port
	fmt.Printf("\033[36m#\033[0m Serving at :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
