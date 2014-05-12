package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	// term colors
	BEGIN = "\033[36m" + "#" + "\033[0m"
)

type Server struct {
	Dir  string
	Port string
	Sub  string
}

func ParseArgs() (*Server, error) {
	dir := flag.String("d", ".", "directory")
	port := flag.String("p", "8080", "port")
	sub := flag.String("s", "/", "server prefix")
	flag.Parse() // handles its own errors
	s := &Server{Dir: *dir, Port: *port, Sub: *sub}
	return s, nil
}

func main() {
	s, err := ParseArgs()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle(s.Sub, http.StripPrefix(s.Sub, http.FileServer(http.Dir(s.Dir))))
	// Serves at port
	fmt.Printf(BEGIN+" Serving at :%s\n", s.Port)
	log.Fatal(http.ListenAndServe(":"+s.Port, nil))
}
