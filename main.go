package main

import (
	"fmt"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", hanlder)
	http.HandleFunc("/get", countHanlder)
	http.ListenAndServe(":8080", nil)
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
}

func countHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, count)
}
