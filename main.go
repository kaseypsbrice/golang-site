package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
    srcDir = "content"
    destDir = "output"
    tmplDir = "templates"
    layoutTmpl = "layout.html"
)

func main() {
    fmt.Println("Setting up webserver...")

    http.Handle("/static/",
        http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    /* Resolves server requests to images and css files. */

    h1 := func (w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("index.html"))
        tmpl.Execute(w, nil)
    }
    http.HandleFunc("/", h1)
    /* Maps a URL pattern to a handler. */

    log.Fatal(http.ListenAndServe(":8000", nil))
    /* log.Fatal will print out the error if one occurs,
     * then exit the process. Remove log.Fatal in Prod.
     * http.ListenAndServe will setup our webserver on port 8000.*/
}