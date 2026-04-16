package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"asciiartweb/asciiart"
	
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Welcome to the Ascii-art page!")
	
	tmpl, err := template.ParseFiles("templates/asciiart.html")
	
	if err != nil {
		fmt.Println("Error!", err)
		return 
	}

	data := struct {
		Text string
		Banner string
	} {
		Text: "hello",
		Banner: "shadow",
	}

	tmpl.Execute(w, data)

}

func AsciiArtWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	result := asciiart.AsciiArt(text, banner)


	tmpl, err := template.ParseFiles("templates/result.html")
	
	if err != nil {
		fmt.Println("Error!", err)
		return 
	}

	data := struct {
		Result string
	} {
		Result: result,
	}

	tmpl.Execute(w, data)
	
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii-art", AsciiArtWeb)
	log.Println("http.//localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))	

}