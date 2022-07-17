package models

import "gorm.io/gorm"

type Feature struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description int32  `json:"description"`
}

func (f *Feature) Insert(db *gorm.DB) error {
	result := db.Create(&f)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (f *Feature) Update(db *gorm.DB) error {
	result := db.Save(&f)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (f *Feature) Select(db *gorm.DB, id int) error {
	result := db.Model(&f).First(&f, id).Scan(&f)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
