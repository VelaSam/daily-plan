package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	web := http.FileServer(http.Dir("./static"))
	http.Handle("/", web)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/byebye", byebyeHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	fmt.Printf("go to http://localhost:8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method "+r.Method+" is not supported", http.StatusNotFound)
		fmt.Fprintf(w, r.Host)
		return
	}

	fmt.Fprintf(w, "!")
}

func byebyeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/byebye" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method "+r.Method+" is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "BYE BYE YOU DID THIS CORRECTLY")
}

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	var p Person

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := p.Name
	address := p.Address

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
