package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home",func(rw http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(rw, "home.html", nil)
	})


	port := ":4444"
	fmt.Println("Executando na porta " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
