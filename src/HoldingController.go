package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

var holdingMux *goji.Mux

func init() {
	holdingMux = goji.SubMux()

	holdingMux.HandleFunc(pat.Get("/"), index)
	holdingMux.HandleFunc(pat.Get("/filter"), filter)
	holdingMux.HandleFunc(pat.Get("/history"), history)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "holdings.index", nil)
}

func filter(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "holdings.filter", nil)
}

func history(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "holdings.history", nil)
}
