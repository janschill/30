package main

import (
	"fmt"
	"log"
	"net/http"

	"30.janschill.de/main/handlers"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
      log.Println(r.URL.Path)
      f(w, r)
  }
}

func main() {
  fs := http.FileServer(http.Dir("assets/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("/", logging(handlers.IndexHandler))
  http.HandleFunc("/u/", logging(handlers.UserHandler))

  fmt.Println("Starting server on localhost:80")
  http.ListenAndServe(":80", nil)
}
