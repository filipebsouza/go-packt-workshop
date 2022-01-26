package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<b>ROOT</b>"))
}

type chapter02Server struct {
	views int
}

func (s *chapter02Server) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	s.views++
	chapter02Json := struct {
		Views int    `json:"messages"`
		Route string `json:"route"`
	}{
		s.views,
		"chapter02",
	}
	jsonBytes, err := json.Marshal(chapter02Json)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(jsonBytes)
}

func handleChapter01(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<b>Chapter 01</b>"))
}

func handleQueryString(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	name, ok := vl["name"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
}

func main() {
	// Used for simplest handles
	http.HandleFunc("/chapter01", handleChapter01)

	// As we've seen before, http.Handler is any struct having a method with this signature
	// ServeHTTP(w http.ResponseWriter, _ *http.Request)
	http.Handle("/", server{})

	chapter02Server := chapter02Server{}
	http.Handle("/chapter02", &chapter02Server)

	http.HandleFunc("/query", handleQueryString)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
