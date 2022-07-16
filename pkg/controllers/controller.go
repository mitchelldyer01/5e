package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

type Controller interface {
	New(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Default struct {
	DB *gorm.DB
}

func (d *Default) New(w http.ResponseWriter, r *http.Request)    {}
func (d *Default) Get(w http.ResponseWriter, r *http.Request)    {}
func (d *Default) Update(w http.ResponseWriter, r *http.Request) {}
