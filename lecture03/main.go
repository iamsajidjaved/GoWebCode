package main

import (
	"log"
	"net/http"
	"text/template"
)

var parsedTemplate = template.Must(template.ParseGlob("*.html"))

type person struct {
	fname string
	lname string
}

func main() {
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/login", handleForm)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Error starting the server", err)
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {

	err := parsedTemplate.ExecuteTemplate(w, "form.html", nil)

	if err != nil {
		log.Println("Error loading the template", err)
		return
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	parse_err := r.ParseForm()
	if parse_err != nil {
		log.Println("Unable to parse the form", parse_err)
		return
	}

	var info person
	info.fname = r.FormValue("fname")
	info.lname = r.FormValue("lname")

	template_err := parsedTemplate.ExecuteTemplate(w, "result.html", struct{ info person }{info})

	if template_err != nil {
		log.Println("Error loading the template", template_err)
		return
	}
}
