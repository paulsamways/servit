package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"
)

var addr = flag.String("addr", ":8000", "address to listen on")
var path = flag.String("path", ".", "path of the static files to serve")
var proxy = flag.String("proxy", "", "sends all requests to [path] to [server], syntax is [path]->[server]")

func main() {
	flag.Parse()

	absPath, err := filepath.Abs(*path)

	if err != nil {
		log.Fatalf("Could not get the absolute path of %v. %v", *path, err)
	}

	if *proxy != "" {
		ps := strings.Split(*proxy, "->")
		u, err := url.Parse(ps[1])
		if err != nil {
			log.Fatal("Couldn't parse proxy URL")
		}

		log.Printf("Redirecting requests to '%v' to '%v'.", ps[0], ps[1])

		http.Handle(ps[0], http.StripPrefix(ps[0], httputil.NewSingleHostReverseProxy(u)))
	}

	http.Handle("/", http.FileServer(http.Dir(absPath)))

	err = http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatalf("Could not serve static files at path %v. %v", absPath, err)
	}
}
