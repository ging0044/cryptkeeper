package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

var dashboardMux *goji.Mux

func init() {
	dashboardMux = goji.SubMux()

	dashboardMux.HandleFunc(pat.Get("/"), dashboard)
	dashboardMux.HandleFunc(pat.Get("/customize"), customize)
}

func dashboard(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dashboard.index", nil)
}

func customize(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dashboard.customize", nil)
}
