package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mitchelldyer01/5e/pkg/middleware"
	"github.com/mitchelldyer01/5e/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CharacterController struct {
	DB     *gorm.DB
	Router *mux.Router
}

func StartCharacterController(DB *gorm.DB, Router *mux.Router) {
	c := &CharacterController{DB: DB, Router: Router}
	c.Router.Use(middleware.Authenticate)
	c.Router.HandleFunc("/character", c.New).Methods("POST")
	c.Router.HandleFunc("/character/{id}", c.Get).Methods("GET")
	c.Router.HandleFunc("/charater", c.Update).Methods("PUT")
	logrus.Println("CharacterController: Initialized \u2705")

	DB.AutoMigrate(&models.Character{})
	logrus.Println("CharacterModel: Seeded \u2705")
}

func (c *CharacterController) New(w http.ResponseWriter, r *http.Request) {
	o := &models.Character{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("Not a valid character: %s", err)
		http.Error(w, "Not a valid character", http.StatusBadRequest)
		return
	}

	err = o.Insert(c.DB)
	if err != nil {
		logrus.Errorf("Error creating character in DB: %s", err)
		http.Error(w, "Error creating character in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (c *CharacterController) Get(w http.ResponseWriter, r *http.Request) {
	o := &models.Character{}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Errorf("Failed parsing ID from string to int: %s", err)
		http.Error(w, "Failed parsing ID from string to int", http.StatusBadRequest)
		return
	}

	err = o.Select(c.DB, id)
	if err != nil {
		logrus.Errorf("Error getting character from DB: %s", err)
		http.Error(w, "Error getting character from DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(o)
}

func (c *CharacterController) Update(w http.ResponseWriter, r *http.Request) {
	o := &models.Character{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("ID is not a valid integer: %s", err)
		http.Error(w, "ID is not a valid integer", http.StatusBadRequest)
		return
	}

	err = o.Select(c.DB, o.ID)
	if err != nil {
		logrus.Errorf("Character not found: %s", err)
		http.Error(w, "Character not found", http.StatusInternalServerError)
		return
	}

	err = o.Update(c.DB)
	if err != nil {
		logrus.Errorf("Error updating character in DB: %s", err)
		http.Error(w, "Error updating character in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
