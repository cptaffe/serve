package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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
	s := new(Server)
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-p" {
			i++
			s.Port = os.Args[i]
		} else if os.Args[i] == "-d" {
			i++
			s.Dir = os.Args[i]
		} else if os.Args[i] == "-s" {
			i++
			s.Sub = os.Args[i]
		} else if os.Args[i] == "-h" || os.Args[i] == "help" {
			fmt.Printf("serve [ -p port | -d dir | -s sub ]\n")
			return s, errors.New("requested help")
		} else {
			s.Port = os.Args[i]
		}
	}
	if s.Port == "" {
		s.Port = "4000"
	}
	if s.Dir == "" {
		s.Dir = "."
	}
	if s.Sub == "" {
		s.Sub = "/"
	}
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
