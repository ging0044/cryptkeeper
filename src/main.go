package main

import (
	"net/http"
	"html/template"

	"goji.io"
	"goji.io/pat"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("src/template/*.gohtml"))
}

func main() {
	//init mux
	mux := goji.NewMux();

	mux.Handle(pat.New("/public/*"), http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	mux.HandleFunc(pat.New("/"), index)
	mux.HandleFunc(pat.New("/login"), login)
	mux.HandleFunc(pat.New("/dashboard"), dashboard)
	mux.HandleFunc(pat.New("/history"), history)
	mux.HandleFunc(pat.New("/filter"), filter)
	mux.HandleFunc(pat.New("/logout"), logout)

	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "login", nil)
}

func dashboard(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dashboard", nil)
}

func history(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "history", nil)
}

func filter(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "filter", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	// log out i guess?
}