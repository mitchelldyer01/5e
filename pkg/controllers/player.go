package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchelldyer01/5e/pkg/middleware"
	"github.com/mitchelldyer01/5e/pkg/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PlayerController struct {
	DB     *gorm.DB
	Router *mux.Router
}

func StartPlayerController(DB *gorm.DB, Router *mux.Router) {
	p := &PlayerController{DB: DB, Router: Router}
	p.Router.Use(middleware.Authenticate)
	p.Router.HandleFunc("/player/register", p.Register).Methods("POST")
	p.Router.HandleFunc("/player/login", p.Login).Methods("POST")
	logrus.Println("PlayerController: Initialized \u2705")

	DB.AutoMigrate(&models.Player{})
	logrus.Println("PlayerModel: Seeded \u2705")
}

func (p *PlayerController) Register(w http.ResponseWriter, r *http.Request) {
	o := &models.Player{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("Invalid registration: %s", err)
		http.Error(w, "Invalid registration", http.StatusBadRequest)
		return
	}

	err = o.Insert(p.DB)
	if err != nil {
		logrus.Errorf("Error registering player in DB: %s", err)
		http.Error(w, "Error registering player in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Registered")
}

func (p *PlayerController) Login(w http.ResponseWriter, r *http.Request) {
	o := &models.Player{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("Invalid login: %s", err)
		http.Error(w, "Invalid login", http.StatusBadRequest)
		return
	}

	err = o.SelectLogin(p.DB)
	if err != nil {
		logrus.Errorf("Login failed: %s", err)
		http.Error(w, "Login failed", http.StatusInternalServerError)
	}

	err = o.UpdateJWT(p.DB)
	if err != nil {
		logrus.Errorf("Login failed: %s", err)
		http.Error(w, "Login failed", http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   o.Token,
		Expires: o.Expiration,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Logged in")
}

func (p *PlayerController) Update(w http.ResponseWriter, r *http.Request) {
	o := &models.Player{}

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&o)
	if err != nil {
		logrus.Errorf("ID is not a valid integer: %s", err)
		http.Error(w, "ID is not a valid integer", http.StatusBadRequest)
		return
	}

	err = o.Update(p.DB)
	if err != nil {
		logrus.Errorf("Error updating player in DB: %s", err)
		http.Error(w, "Error updating player in DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
