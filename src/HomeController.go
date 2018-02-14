package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

var homeMux *goji.Mux

func init() {
	homeMux = goji.SubMux()

	homeMux.HandleFunc(pat.Get("/"), home)
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "home.index", "This is some content")
}
