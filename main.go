package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageData struct {
	AdelaideTime string
}

func getAdelaideTime() (string, string) {
	location, err := time.LoadLocation("Australia/Adelaide")
	if err != nil {
		log.Println("Error loading Adelaide time zone:", err)
		return "", ""
	}
	adelaideTime := time.Now().In(location)
	hours := adelaideTime.Format("15")  // 24-hour format for hours
	minutes := adelaideTime.Format("04") // Minutes without a leading zero
	return hours, minutes
	// 15:04 is a reference for how the time should be displayed
	// (in 24 hour time)
}

func handleTemplates(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"index.html",
		"templates/navbar.html",
		"templates/home.html",
	))

	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func timeAPI(w http.ResponseWriter, r *http.Request) {
	hours, minutes := getAdelaideTime()
	response := map[string]string{
		"hours":   hours,
		"minutes": minutes,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
// Returns current Adelaide time as JSON

func main() {
	fmt.Println("Starting server on :8000...")

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	/* Resolves server requests to static files e.g. images, css files. */

	http.HandleFunc("/", handleTemplates)
	http.HandleFunc("/api/time", timeAPI)
	/* Maps routes / endpoints */

	log.Fatal(http.ListenAndServe(":8000", nil))
	/* log.Fatal will print out the error if one occurs,
	 * then exit the process. Remove log.Fatal in Prod.
	 * http.ListenAndServe will setup our webserver on port 8000.*/
}
