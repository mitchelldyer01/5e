package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mitchelldyer01/5e/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SpellController struct {
	DB     *gorm.DB
	Router *mux.Router
}

func NewSpellController(DB *gorm.DB, Router *mux.Router) {
	c := &SpellController{DB: DB, Router: Router}
	c.Router.HandleFunc("/spell", c.New).Methods("POST")
	c.Router.HandleFunc("/spell/{id}", c.Get).Methods("GET")
	c.Router.HandleFunc("/spell", c.Update).Methods("PUT")
	logrus.Println("SpellController: Initialized \u2705")

	DB.AutoMigrate(&models.Spell{})
	logrus.Println("SpellModel: Seeded \u2705")
}

func (s *SpellController) New(w http.ResponseWriter, r *http.Request) {
	o := &models.Spell{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("Not a valid spell: %s", err)
		http.Error(w, "Not a valid spell", http.StatusBadRequest)
		return
	}

	err = o.Insert(s.DB)
	if err != nil {
		logrus.Errorf("Error creating spell in DB: %s", err)
		http.Error(w, "Error creating spell in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (s *SpellController) Get(w http.ResponseWriter, r *http.Request) {
	o := &models.Spell{}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Errorf("Failed parsing ID from string to int: %s", err)
		http.Error(w, "Failed parsing ID from string to int", http.StatusBadRequest)
		return
	}

	err = o.Select(s.DB, id)
	if err != nil {
		logrus.Errorf("Error getting spell from DB: %s", err)
		http.Error(w, "Error getting spell from DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(o)
}

func (s *SpellController) Update(w http.ResponseWriter, r *http.Request) {
	o := &models.Spell{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("ID is not a valid integer: %s", err)
		http.Error(w, "ID is not a valid integer", http.StatusBadRequest)
		return
	}

	err = o.Select(s.DB, o.ID)
	if err != nil {
		logrus.Errorf("Spell not found: %s", err)
		http.Error(w, "Spell not found", http.StatusInternalServerError)
		return
	}

	err = o.Update(s.DB)
	if err != nil {
		logrus.Errorf("Error updating spell in DB: %s", err)
		http.Error(w, "Error updating spell in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
