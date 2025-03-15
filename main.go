package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleTemplates(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"index.html",
		"templates/navbar.html",
		"templates/home.html",
	))
	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("Starting server on :8000...")

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	/* Resolves server requests to static files e.g. images, css files. */

	http.HandleFunc("/", handleTemplates)
	/* Maps a URL pattern to a handler. */

	log.Fatal(http.ListenAndServe(":8000", nil))
	/* log.Fatal will print out the error if one occurs,
	 * then exit the process. Remove log.Fatal in Prod.
	 * http.ListenAndServe will setup our webserver on port 8000.*/
}
