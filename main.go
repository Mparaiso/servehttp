package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		port   = flag.String("port", "80", "port the server will be listening on")
		ip     = flag.String("id", "0.0.0.0", "server ip")
		root   = flag.String("root", "", "server root folder")
		addr   string
		err    error
		server *http.ServeMux
	)
	flag.Parse()
	if *root == "" {
		*root, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	addr = *ip + ":" + *port
	server = http.NewServeMux()
	fileHandler := http.FileServer(http.Dir(*root))
	logger := func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		fileHandler.ServeHTTP(rw, r)
	}
	server.HandleFunc("/", logger)
	log.Println("Listening on:", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}
