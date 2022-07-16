package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchelldyer01/characters-5e/pkg/controllers"
	"github.com/mitchelldyer01/characters-5e/pkg/db"
)

func main() {
	repo := db.New()

	c := controllers.CharacterController{DB: repo.DB}

	r := mux.NewRouter()
	r.HandleFunc("/characterl", c.New).Methods("POST")
	r.HandleFunc("/character/{id}", c.Get).Methods("GET")
	r.HandleFunc("/charater", c.Update).Methods("PUT")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
