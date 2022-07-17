package models

import "gorm.io/gorm"

type Action struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description int32  `json:"description"`
}

func (a *Action) Insert(db *gorm.DB) error {
	result := db.Create(&a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *Action) Update(db *gorm.DB) error {
	result := db.Save(&a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *Action) Select(db *gorm.DB, id int) error {
	result := db.Model(&a).First(&a, id).Scan(&a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
