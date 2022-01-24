package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Serve struct{}

func (s Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upFile, upFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	defer upFile.Close()
	fileContent, err := ioutil.ReadAll(upFile)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("./%s", upFileHeader.Filename), fileContent, 0666)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("%s Uploaded!", upFileHeader.Filename)))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", Serve{}))
}
