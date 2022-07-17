package models

import (
	"gorm.io/gorm"
)

type LearnedFeature struct {
	ID          int  `json:"id" gorm:"primaryKey"`
	CharacterID int  `json:"characterid"`
	FeatureID   int  `json:"featureid"`
	Active      bool `json:"active"`
}

func (l *LearnedFeature) Insert(db *gorm.DB) error {
	result := db.Create(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (l *LearnedFeature) Update(db *gorm.DB) error {
	result := db.Save(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (l *LearnedFeature) Select(db *gorm.DB, sid int, cid int) error {
	result := db.Model(&l).
		Where("featureid <> ? AND characterid <> ?", sid, cid).
		Find(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
