package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	// colors (linux)
	BLUE  = "\033[36m"
	RED   = "\033[31m"
	ENDC  = "\033[0m"
	BEGIN = BLUE + "#" + ENDC
	ERROR = "(" + RED + "!" + ENDC + ")"
)

type Server struct {
	Dir  string
	Port string
	Sub  string
}

func ParseArgs() (*Server, error) {
	dir := flag.String("d", ".", "directory")
	port := flag.String("p", "4000", "port")
	sub := flag.String("s", "/", "server prefix")
	flag.Parse()
	s := &Server{Dir: *dir, Port: *port, Sub: *sub}
	return s, nil
}

func main() {
	s, err := ParseArgs()
	if err != nil {
		fmt.Printf(BEGIN+" %s "+ERROR+"\n", err)
		return
	}

	http.Handle(s.Sub, http.StripPrefix(s.Sub, http.FileServer(http.Dir(s.Dir))))
	// Serves at port
	fmt.Printf("\033[36m#\033[0m Serving at :%s\n", s.Port)
	log.Fatal(http.ListenAndServe(":"+s.Port, nil))
}
