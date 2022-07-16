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

type LearnedController struct {
	DB     *gorm.DB
	Router *mux.Router
}

func NewLearnedController(DB *gorm.DB, Router *mux.Router) {
	l := &LearnedController{DB: DB, Router: Router}
	l.Router.HandleFunc("/learned", l.New).Methods("POST")
	l.Router.HandleFunc("/learned/spell/{sid}/character/{cid}", l.Get).Methods("GET")
	l.Router.HandleFunc("/learned", l.Update).Methods("PUT")
	logrus.Println("LearnedController: Initialized \u2705")

	DB.AutoMigrate(&models.Learned{})
	logrus.Println("LearnedModel: Seeded \u2705")
}

func (l *LearnedController) New(w http.ResponseWriter, r *http.Request) {
	o := &models.Learned{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = o.Insert(l.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (l *LearnedController) Get(w http.ResponseWriter, r *http.Request) {
	o := &models.Learned{}

	vars := mux.Vars(r)
	sid, err := strconv.Atoi(vars["sid"])
	if err != nil {
		logrus.Errorf("Failed parsing SID from string to int: %s", err)
		http.Error(w, "Failed parsing SID from string to int", http.StatusBadRequest)
		return
	}
	cid, err := strconv.Atoi(vars["cid"])
	if err != nil {
		logrus.Errorf("Failed parsing CID from string to int: %s", err)
		http.Error(w, "Failed parsing CID from string to int", http.StatusBadRequest)
		return
	}

	err = o.Select(l.DB, sid, cid)
	if err != nil {
		logrus.Errorf("Error getting spell from DB: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(o)
}

func (l *LearnedController) Update(w http.ResponseWriter, r *http.Request) {
	o := &models.Learned{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("ID is not a valid integer: %s", err)
		http.Error(w, "ID is not a valid integer", http.StatusBadRequest)
		return
	}

	err = o.Update(l.DB)
	if err != nil {
		logrus.Errorf("Error updating spell in DB: %s", err)
		http.Error(w, "Error updating spell in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
