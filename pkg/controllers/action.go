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

type ActionController struct {
	DB     *gorm.DB
	Router *mux.Router
}

func StartActionController(DB *gorm.DB, Router *mux.Router) {
	s := &ActionController{DB: DB, Router: Router}
	s.Router.Use(middleware.Authenticate)
	s.Router.HandleFunc("/action", s.New).Methods("POST")
	s.Router.HandleFunc("/action/{id}", s.Get).Methods("GET")
	s.Router.HandleFunc("/action", s.Update).Methods("PUT")
	logrus.Println("ActionController: Initialized \u2705")

	DB.AutoMigrate(&models.Action{})
	logrus.Println("ActionModel: Seeded \u2705")
}

func (a *ActionController) New(w http.ResponseWriter, r *http.Request) {
	o := &models.Action{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("Not a valid action: %s", err)
		http.Error(w, "Not a valid action", http.StatusBadRequest)
		return
	}

	err = o.Insert(a.DB)
	if err != nil {
		logrus.Errorf("Error creating action in DB: %s", err)
		http.Error(w, "Error creating action in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (a *ActionController) Get(w http.ResponseWriter, r *http.Request) {
	o := &models.Action{}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Errorf("Failed parsing ID from string to int: %s", err)
		http.Error(w, "Failed parsing ID from string to int", http.StatusBadRequest)
		return
	}

	err = o.Select(a.DB, id)
	if err != nil {
		logrus.Errorf("Error getting action from DB: %s", err)
		http.Error(w, "Error getting action from DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(o)
}

func (a *ActionController) Update(w http.ResponseWriter, r *http.Request) {
	o := &models.Spell{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("ID is not a valid integer: %s", err)
		http.Error(w, "ID is not a valid integer", http.StatusBadRequest)
		return
	}

	err = o.Select(a.DB, o.ID)
	if err != nil {
		logrus.Errorf("Action not found: %s", err)
		http.Error(w, "Action not found", http.StatusInternalServerError)
		return
	}

	err = o.Update(a.DB)
	if err != nil {
		logrus.Errorf("Error updating action in DB: %s", err)
		http.Error(w, "Error updating action in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
