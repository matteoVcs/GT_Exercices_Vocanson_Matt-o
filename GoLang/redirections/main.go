package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Redirect struct {
	Redirect1 string
	Redirect2 string
	Redirect3 string
}

func main() {
	var g Redirect
	fmt.Println("server is running on port 8080")
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", g.Index)
	http.HandleFunc("/page1", g.Page1)
	http.HandleFunc("/page2", g.Page2)
	http.HandleFunc("/page3", g.Page3)
	http.ListenAndServe(":8080", nil)

}

func (g *Redirect) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp2 := template.Must(template.ParseFiles("index.html"))
	details := Redirect{
		Redirect1: r.FormValue("redirect1"),
		Redirect2: r.FormValue("redirect2"),
		Redirect3: r.FormValue("redirect3"),
	}
	g.Testform(r, w)
	tmp2.Execute(w, details)
}

func (g *Redirect) Testform(r *http.Request, w http.ResponseWriter) {
	if r.FormValue("redirect") == "page1" {
		http.Redirect(w, r, "/page1", 302)
		return
	}
	if r.FormValue("redirect") == "page2" {
		http.Redirect(w, r, "/page2", 302)
		return
	}
	if r.FormValue("redirect") == "page3" {
		http.Redirect(w, r, "/page3", 302)
		return
	}
}

func (g *Redirect) Page1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page1" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp2 := template.Must(template.ParseFiles("page1.html"))
	details := Redirect{}
	tmp2.Execute(w, details)
}

func (g *Redirect) Page2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page2" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp2 := template.Must(template.ParseFiles("page2.html"))
	details := Redirect{}
	tmp2.Execute(w, details)
}

func (g *Redirect) Page3(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/page3" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmp2 := template.Must(template.ParseFiles("page3.html"))
	details := Redirect{}
	tmp2.Execute(w, details)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 not found")
	}
}
