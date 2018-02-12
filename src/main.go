package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	//init mux
	mux := goji.NewMux();

	mux.Handle(pat.Get("/public/*"), http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	mux.Handle(pat.Get("/user/*"), userMux)
	mux.Handle(pat.Get("/dashboard/*"), dashboardMux)
	mux.Handle(pat.Get("/holdings/*"), holdingMux)
	mux.Handle(pat.Get("/*"), homeMux)

	http.ListenAndServe(":8080", mux)
}
