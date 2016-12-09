package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8070", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
