package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type StaticFiles struct {
	prefix string
	dir    http.Dir
}
type Image struct {
	ArrayBytes []byte
}

type PageData struct {
	Title string
}

func (s StaticFiles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(s.dir)
	http.StripPrefix(s.prefix, fs).ServeHTTP(w, r)
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, erro := template.ParseFiles("./index.html")
	if erro != nil {
		http.Error(w, "Not Found HTML", http.StatusBadRequest)
	}
	pageData := PageData{
		Title: "Anna Reis",
	}
	tmpl.Execute(w, pageData)
}

func main() {
	fmt.Println("Server running on PORT 8080")
	http.Handle("/static/", StaticFiles{prefix: "/static/", dir: http.Dir("./static/")})
	http.HandleFunc("/", serveRoot)
	http.ListenAndServe(":8080", nil)
}
