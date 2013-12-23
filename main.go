package main 

import (
	"log"
	"flag"
	"net/http"
)

var (
	ListenAddr = flag.String("addr", ":80", "The ADDRESS:PORT to listen on")
	DocumentRoot = flag.String("dir", "", "The directory to serve up")
)

func main() {
	flag.Parse()
	if *DocumentRoot == "" {
		log.Fatalln("You must specify a directory to serve, with '-dir=\"...\"'")
	}

	handler := http.FileServer(http.Dir(*DocumentRoot))
	log.Printf("Serving %q", *DocumentRoot)
	log.Printf("Listening on %q", *ListenAddr)
	if err := http.ListenAndServe(*ListenAddr, handler); err != nil {
		log.Fatalln(err)
	}

	return
}
