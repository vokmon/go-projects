package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
	}

	fmt.Fprintf(w, "Post request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)
}

func main() {
	port := 8080

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Start server at port %v\n", port)

	var portString = fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal(err)
	}
}
