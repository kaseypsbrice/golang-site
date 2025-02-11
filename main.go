package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
    fmt.Println("Setting up webserver...")

    h1 := func (w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("index.html"))
        tmpl.Execute(w, nil)
    }
    http.HandleFunc("/", h1)
    /* Maps a URL pattern to a handler. */

    log.Fatal(http.ListenAndServe(":8000", nil))
    /* log.Fatal will print out the error if one occurs,
     * then exit the process. 
     * http.ListenAndServe will setup our webserver on port 8000.*/
}
