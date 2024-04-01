package main 

import (
	"html/template"
	"net/http"
)

type Project struct {
    Title string
    Done  bool
}

type ProjectCentral struct {
	Titre string
	Projets []Project
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/",func(w http.ResponseWriter, r*http.Request) {
		data := ProjectCentral{
			Titre : "Gestionnaire",
			Projets : []Project{
				{Title: "BlueBird", Done: true },
				{Title: "RedFox", Done: false},
				{Title: "YellowShark", Done: false},
			},
		}
		tmpl.Execute(w,data)
	})
	http.ListenAndServe("localhost:80", nil)
}
