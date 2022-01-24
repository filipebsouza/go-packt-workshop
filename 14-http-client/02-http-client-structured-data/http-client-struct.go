package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

type server struct{}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "{ \"message\": \"hello world.\" }"
	w.Write([]byte(msg))
}

func getDataAndReturnResponse() messageData {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	go http.ListenAndServe(":8080", server{})
	go getDataAndReturnResponse()
}
