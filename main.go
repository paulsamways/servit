package main

import (
  "net/http"
  "log"
  "flag"
  "path/filepath"
)

var addr = flag.String("addr", ":8000", "address to listen on")
var path = flag.String("path", ".", "path of the static files to serve")
var noCache = flag.Bool("noCache", true, "sends the Cache-Control headers to the client to prevent caching")

func main() {
  flag.Parse()

  absPath, err := filepath.Abs(*path)

  if err != nil {
    log.Fatalf("Could not get the absolute path of %v. %v", *path, err)
  }

  h := http.FileServer(http.Dir(absPath))

  if *noCache {
    h = CachePreventionHandler(h)
  }

  err = http.ListenAndServe(*addr, h)

  if err != nil {
    log.Fatalf("Could not serve static files at path %v. %v", absPath, err)
  }
}

func CachePreventionHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "max-age=0, no-cache, no-store")
    h.ServeHTTP(w, r)
  })
}
