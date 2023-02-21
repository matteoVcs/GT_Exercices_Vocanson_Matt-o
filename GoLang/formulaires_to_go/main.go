package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Global struct {
	Text string
}

func main() {
	var g Global
	fmt.Println("server is running on port 8080")
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", g.Formulaire)
	http.ListenAndServe(":8080", nil)

}

func (g *Global) Formulaire(w http.ResponseWriter, r *http.Request) {
	tmp2 := template.Must(template.ParseFiles("index.html"))
	details := Global{
		Text: r.FormValue("letter"),
	}
	tmp2.Execute(w, details)
	Testform(r, w)
}

func Testform(r *http.Request, w http.ResponseWriter) {
	fmt.Println(r.FormValue("letter"))
}
