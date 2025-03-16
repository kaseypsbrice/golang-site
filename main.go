package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Project struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
}

func loadProjects() ([]Project, error) {
	data, err := os.ReadFile("projects.json")
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(data, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func getAdelaideTime() (string, string) {
	location, err := time.LoadLocation("Australia/Adelaide")
	if err != nil {
		log.Println("Error loading Adelaide time zone:", err)
		return "", ""
	}
	adelaideTime := time.Now().In(location)
	hours := adelaideTime.Format("15")
	minutes := adelaideTime.Format("04")
	return hours, minutes
	// 15:04 is a reference for how the time should be displayed
	// (in 24 hour time)
}

func handleTemplates(w http.ResponseWriter, r *http.Request) {
	projects, err := loadProjects()
	if err != nil {
		http.Error(w, "Failed to load projects", http.StatusInternalServerError)
		return
	}

	data:= struct {
		Projects []Project
	}{
		Projects: projects,
	}

	tmpl := template.Must(template.ParseFiles(
		"index.html",
		"templates/navbar.html",
		"templates/home.html",
		"templates/projects.html",
		"templates/sticky_note.html",
	))

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func timeAPI(w http.ResponseWriter, r *http.Request) {
	hours, minutes := getAdelaideTime()

	response := fmt.Sprintf(`
		<div class="flex flex-col">
			<span class="text-4xl">%s</span>
			<span class="text-4xl">%s</span>
			<span class="text-1xl">ACST</span>
		</div>
	`, hours, minutes)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}
// Returns current Adelaide time in HTML

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
