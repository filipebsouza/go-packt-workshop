package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type server struct{}
type messageData struct {
	Message string `json:"message"`
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	messageData := messageData{}
	err := jsonDecoder.Decode(&messageData)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, _ := json.Marshal(messageData)
	log.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func postDataAndReturnResponse(msg messageData) messageData {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	r, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
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
	go func() {
		log.Fatal(http.ListenAndServe(":8080", server{}))
	}()
	go func() {
		postDataAndReturnResponse(messageData{Message: "Hi server"})
	}()

	time.Sleep(600000)
}
