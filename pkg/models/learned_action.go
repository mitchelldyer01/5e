package models

import (
	"gorm.io/gorm"
)

type LearnedAction struct {
	ID          int `json:"id" gorm:"primaryKey"`
	CharacterID int `json:"characterid"`
	ActionID    int `json:"actionid"`
}

func (l *LearnedAction) Insert(db *gorm.DB) error {
	result := db.Create(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (l *LearnedAction) Update(db *gorm.DB) error {
	result := db.Save(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (l *LearnedAction) Select(db *gorm.DB, sid int, cid int) error {
	result := db.Model(&l).
		Where("actionid <> ? AND characterid <> ?", sid, cid).
		Find(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
