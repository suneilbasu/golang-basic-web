package main

import (
	"fmt"
	"net/http"
)

var names = []string{
	"George",
	"Suneil",
}

func handler(w http.ResponseWriter, r *http.Request) {
	for i, name := range names {
		fmt.Fprintf(w, "%d: %s\n", i, name)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Trying to run")
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
