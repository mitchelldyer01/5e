package models

import (
	"gorm.io/gorm"
)

type LearnedSpell struct {
	ID          int  `json:"id" gorm:"primaryKey"`
	CharacterID int  `json:"characterid"`
	SpellID     int  `json:"spellid"`
	Active      bool `json:"active"`
}

func (l *LearnedSpell) Insert(db *gorm.DB) error {
	result := db.Create(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (l *LearnedSpell) Update(db *gorm.DB) error {
	result := db.Save(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (l *LearnedSpell) Select(db *gorm.DB, sid int, cid int) error {
	result := db.Model(&l).
		Where("spellid <> ? AND characterid <> ?", sid, cid).
		Find(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
