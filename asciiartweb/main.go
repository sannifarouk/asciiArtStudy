package main

import (
	"log"
	"net/http"
	"html/template"
	"asciiartweb/asciiart"
	
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Welcome to the Ascii-art page!")
	
	tmpl, err := template.ParseFiles("templates/asciiart.html")
	
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound) // 404
		return 
	}

	tmpl.Execute(w, nil)

}

func AsciiArtWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	
	if text == "" {
		http.Error(w, "Empty text submitted", http.StatusBadRequest) // 400
		return
	}

	if banner == "" {
		http.Error(w, "No banner selected", http.StatusBadRequest) // 400
		return
	}

	result, err := asciiart.AsciiArt(text, banner)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound) // 404
		return
	}

	tmpl, err := template.ParseFiles("templates/asciiart.html")
	if err != nil {
		// fmt.Println("Error!", err)
		http.Error(w, "Template Error", http.StatusInternalServerError) // 500
		return 
	}

	w.WriteHeader(http.StatusOK) // 200
	// fmt.Fprintln(w, "Success")

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