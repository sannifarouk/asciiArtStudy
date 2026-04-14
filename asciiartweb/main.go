package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Ascii-art page!")

}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii-art", AsciiArt)
	log.Println("http.//localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}