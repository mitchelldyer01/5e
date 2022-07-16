package models

import "gorm.io/gorm"

type Character struct {
	ID                  int    `json:"id" gorm:"primaryKey"`
	Acrobatics          int32  `json:"acrobatics"`
	AnimalHandling      int32  `json:"animalhandling"`
	Arcana              int32  `json:"arcana"`
	ArmorClass          int32  `json:"armorclass"`
	Athletics           int32  `json:"athletics"`
	Class               string `json:"class"`
	Charisma            int32  `json:"charisma"`
	Constitution        int32  `json:"constitution"`
	Darkvision          bool   `json:"darkvision"`
	Dexterity           int32  `json:"dexterity"`
	History             int32  `json:"history"`
	HP                  int32  `json:"hp"`
	Initiative          int32  `json:"initiative"`
	Insight             int32  `json:"insight"`
	Intelligence        int32  `json:"intelligence"`
	Intimidation        int32  `json:"intimidation"`
	Investigation       int32  `json:"investigation"`
	Level               int32  `json:"level"`
	Medicine            int32  `json:"medicine"`
	Name                string `json:"name"`
	Nature              int32  `json:"nature"`
	Perception          int32  `json:"perception"`
	Performance         int32  `json:"performance"`
	Persuasion          int32  `json:"persuasion"`
	Proficiency         int32  `json:"proficiency"`
	Race                string `json:"race"`
	Religion            int32  `json:"religion"`
	SleightOfHand       int32  `json:"sleightofhand"`
	SpellSlotsOne       int32  `json:"spellslotsone"`
	SpellSlotsTwo       int32  `json:"spellslotstwo"`
	SpellSlotsThree     int32  `json:"spellslotsthree"`
	SpellSlotsFour      int32  `json:"spellslotsfour"`
	SpellSlotsFive      int32  `json:"spellslotsfive"`
	SpellSlotsSix       int32  `json:"spellslotssix"`
	SpellSlotsSeven     int32  `json:"spellslotsseven"`
	SpellSlotsEight     int32  `json:"spellslotseight"`
	SpellSlotsNine      int32  `json:"spellslotsnine"`
	Stealth             int32  `json:"stealth"`
	Strength            int32  `json:"strength"`
	Survival            int32  `json:"survival"`
	Subclass            string `json:"subclass"`
	Subrace             string `json:"subrace"`
	UsedSpellSlotsOne   int32  `json:"usedspellslotsone"`
	UsedSpellSlotsTwo   int32  `json:"usedspellslotstwo"`
	UsedSpellSlotsThree int32  `json:"usedspellslotsthree"`
	UsedSpellSlotsFour  int32  `json:"usedspellslotsfour"`
	UsedSpellSlotsFive  int32  `json:"usedspellslotsfive"`
	UsedSpellSlotsSix   int32  `json:"usedspellslotssix"`
	UsedSpellSlotsSeven int32  `json:"usedspellslotsseven"`
	UsedSpellSlotsEight int32  `json:"usedspellslotseight"`
	UsedSpellSlotsNine  int32  `json:"usedspellslotsnine"`
	Vision              int32  `json:"vision"`
	Wisdom              int32  `json:"wisdom"`
}

func (c *Character) Insert(db *gorm.DB) error {
	result := db.Create(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Character) Update(db *gorm.DB) error {
	result := db.Save(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Character) Select(db *gorm.DB, id int) error {
	result := db.Model(&c).First(&c, id).Scan(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
