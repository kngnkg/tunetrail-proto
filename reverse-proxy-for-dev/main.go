package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	log.Print("Starting reverse proxy server on port 8080")
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		// request.URL.Host = "docker.for.mac.localhost:18000"
		request.URL.Host = "tunetrail-restapi:8080"
	}

	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":443",
		Handler: rp,
	}

	if err := server.ListenAndServeTLS("localhost.pem","localhost-key.pem"); err != nil {
		log.Fatal(err.Error())
	}
}
