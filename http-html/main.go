package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template


type Usuario struct {
	Name string
	Idade int
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home",func(rw http.ResponseWriter, r *http.Request) {

		u := Usuario{
			Name: "Example",
			Idade: 22,
		}

		templates.ExecuteTemplate(rw, "home.html", u)
	})


	port := ":4444"
	fmt.Println("Executando na porta " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
