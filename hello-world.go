package main

import (
	"net/http"
)

func main() {

	//ROUTES
	http.HandleFunc("/", homeHandle)
	http.HandleFunc("/contact", contactHandle)

	//STAR SERVER
	http.ListenAndServe(":3000", nil)
}

func homeHandle(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello world"))
}

func contactHandle(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("contact page"))
}