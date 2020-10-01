package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ttprinter(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("timetable90.txt")
	if err != nil {
		l := log.New(os.Stderr, "", 1)
		l.Println(err)
		os.Exit(1)
	}
	lines := string(b)
	io.WriteString(w, lines)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", ttprinter)
	http.ListenAndServe(":"+port, nil)
}
