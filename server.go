package main

import (
	// "fmt"

	"github.com/dcao96/Simple-Go-Key-Server/routes"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	// setup routers
	r := mux.NewRouter()
	// Use this for subdomains
	// s := r.Host("localhost").Subrouter()

	r.HandleFunc("/register", routes.RegisterHandler).Methods("POST")
	r.HandleFunc("/get", routes.GetHandler).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":8080")
}
