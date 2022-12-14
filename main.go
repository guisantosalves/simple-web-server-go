package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer) // -> get the file server that I set in the var

	// route, callback
	// w -> response (write the response)
	// r -> request (requesting data)
	http.HandleFunc("/hello", handleHello)

	http.HandleFunc("/form", formHandler)

	fmt.Println("Server running at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// response, request
func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "POST request successful")
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "adress = %s\n", address)
}

func handleHello(w http.ResponseWriter, r *http.Request) {

	// handle with errors
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not suported.", http.StatusNotFound)
		return
	}

	// sending a response
	fmt.Fprintf(w, "Hello, boy")
}
