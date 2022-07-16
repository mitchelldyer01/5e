package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Learned struct {
	ID          int  `json:"id" gorm:"primaryKey"`
	CharacterID int  `json:"characterid"`
	SpellID     int  `json:"spellid"`
	Active      bool `json:"active"`
}

func (l *Learned) Insert(db *gorm.DB) error {
	result := db.Create(&l)
	if result.Error != nil {
		logrus.Errorf("Error creating spell in DB: %s", result.Error)
		return result.Error
	}
	return nil
}
func (l *Learned) Update(db *gorm.DB) error {
	result := db.Save(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (l *Learned) Select(db *gorm.DB, sid int, cid int) error {
	result := db.Model(&l).
		Where("spellid <> ? AND characterid <> ?", sid, cid).
		Find(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
