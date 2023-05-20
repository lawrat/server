package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const addr = "localhost"
const port = ":7000"

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)

	fmt.Printf("Serveur en cours d'execution sur : http://%s%s", addr, port)
	http.ListenAndServe(port, nil)
}
