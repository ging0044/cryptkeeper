package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

var userMux *goji.Mux

func init() {
	userMux = goji.SubMux()

	userMux.HandleFunc(pat.Get("/login"), login)
	userMux.HandleFunc(pat.Get("/logout"), logout)
	userMux.HandleFunc(pat.Get("/settings"), settings)
	userMux.HandleFunc(pat.Get("/register"), register)
}

func login(w http.ResponseWriter, req *http.Request) {
	//TODO: log in logic
	tpl.ExecuteTemplate(w, "user.login", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	//TODO: log out logic
}

func settings(w http.ResponseWriter, req *http.Request) {
	//TODO: user settings, and page to change
	tpl.ExecuteTemplate(w, "user.settings", nil)
}

func register(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "user.register", nil)
}
