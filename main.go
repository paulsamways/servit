package main

import (
  "net/http"
  "log"
  "flag"
  "path/filepath"
)

var addr = flag.String("addr", ":8000", "address to listen on")
var path = flag.String("path", ".", "path of the static files to serve")

func main() {
  absPath, err := filepath.Abs(*path)

  if err != nil {
    log.Fatalf("Could not get the absolute path of %v. %v", *path, err)
  }

  err = http.ListenAndServe(*addr, http.FileServer(http.Dir(absPath)))

  if err != nil {
    log.Fatalf("Could not serve static files at path %v. %v", absPath, err)
  }
}
