package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
	"os"
	"log"
	"github.com/unrolled/logger"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		log.Printf("No PORT environment variable set, using 8080...")
		port = "8080"
	}
}

func main() {
	loggerMiddleware := logger.New()

	//init mux
	mux := goji.NewMux();
	mux.Use(loggerMiddleware.Handler)

	mux.Handle(pat.Get("/public/*"), http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	mux.Handle(pat.Get("/user/*"), userMux)
	mux.Handle(pat.Get("/dashboard/*"), dashboardMux)
	mux.Handle(pat.Get("/holdings/*"), holdingMux)
	mux.Handle(pat.Get("/*"), homeMux)

	log.Printf("Listening on port %v", port)
	http.ListenAndServe(":"+port, mux)
}
