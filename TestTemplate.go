package main 

import (
	"fmt"
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

type User struct {
	Titre string
	Pass string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
}

func login(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	data := User{
		Titre: r.FormValue("username"),
		Pass : r.FormValue("password"),
	}
	
	fmt.Println(data)
	fmt.Println(data.Pass)
	// fmt.Println(r.Form["password"])
	template.Must(template.ParseFiles("info.html")).Execute(w,data)
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
	http.HandleFunc("/login",login)
	http.ListenAndServe("localhost:8080", nil)
}
