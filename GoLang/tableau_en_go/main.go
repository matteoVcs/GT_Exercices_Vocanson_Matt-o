package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type Global struct {
	N       int
	Contenu []int
}

func main() {
	var g Global
	fmt.Println("server is running on port 8080")
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", g.Page)
	http.ListenAndServe(":8080", nil)
}

func (g *Global) InitStruct(r *http.Request, w http.ResponseWriter) []int {
	var clear []int
	g.Contenu = clear
	var max int = 100
	var min int = 3
	rand.Seed(time.Now().UnixNano())
	g.N = rand.Intn(max-min) + min
	for i := 0; i < g.N; i++ {
		g.Contenu = append(g.Contenu, rand.Int())
	}
	return g.Contenu
}

func (g *Global) Page(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := Global{
		Contenu: g.InitStruct(r, w),
	}
	tmp.Execute(w, details)
}
