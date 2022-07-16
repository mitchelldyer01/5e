package models

import (
	"gorm.io/gorm"
)

type Spell struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Level       int32  `json:"level"`
	Range       int32  `json:"range"`
	Time        int32  `json:"time"`
	Duration    int32  `json:"duration"`
	Description int32  `json:"description"`
}

func (s *Spell) Insert(db *gorm.DB) error {
	result := db.Create(&s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Spell) Update(db *gorm.DB) error {
	result := db.Save(&s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Spell) Select(db *gorm.DB, id int) error {
	result := db.Model(&s).First(&s, id).Scan(&s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
