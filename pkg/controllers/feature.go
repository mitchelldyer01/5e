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

type FeatureController struct {
	DB     *gorm.DB
	Router *mux.Router
}

func StartFeatureController(DB *gorm.DB, Router *mux.Router) {
	s := &FeatureController{DB: DB, Router: Router}
	s.Router.Use(middleware.Authenticate)
	s.Router.HandleFunc("/feature", s.New).Methods("POST")
	s.Router.HandleFunc("/feature/{id}", s.Get).Methods("GET")
	s.Router.HandleFunc("/feature", s.Update).Methods("PUT")
	logrus.Println("FeatureController: Initialized \u2705")

	DB.AutoMigrate(&models.Feature{})
	logrus.Println("FeatureModel: Seeded \u2705")
}

func (f *FeatureController) New(w http.ResponseWriter, r *http.Request) {
	o := &models.Action{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("Not a valid feature: %s", err)
		http.Error(w, "Not a valid feature", http.StatusBadRequest)
		return
	}

	err = o.Insert(f.DB)
	if err != nil {
		logrus.Errorf("Error creating feature in DB: %s", err)
		http.Error(w, "Error creating feature in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (f *FeatureController) Get(w http.ResponseWriter, r *http.Request) {
	o := &models.Action{}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Errorf("Failed parsing ID from string to int: %s", err)
		http.Error(w, "Failed parsing ID from string to int", http.StatusBadRequest)
		return
	}

	err = o.Select(f.DB, id)
	if err != nil {
		logrus.Errorf("Error getting feature from DB: %s", err)
		http.Error(w, "Error getting feature from DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(o)
}

func (f *FeatureController) Update(w http.ResponseWriter, r *http.Request) {
	o := &models.Spell{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("ID is not a valid integer: %s", err)
		http.Error(w, "ID is not a valid integer", http.StatusBadRequest)
		return
	}

	err = o.Select(f.DB, o.ID)
	if err != nil {
		logrus.Errorf("Feature not found: %s", err)
		http.Error(w, "Feature not found", http.StatusInternalServerError)
		return
	}

	err = o.Update(f.DB)
	if err != nil {
		logrus.Errorf("Error updating feature in DB: %s", err)
		http.Error(w, "Error updating feature in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
