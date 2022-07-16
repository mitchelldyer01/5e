package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestString string = "Test"
	TestInt    int32  = 1
)

// Test setting & getting all of the properties
func Test_NewChar(t *testing.T) {
	c := &Character{}

	c.Acrobatics = TestInt
	c.AnimalHandling = TestInt
	c.Arcana = TestInt
	c.ArmorClass = TestInt
	c.Athletics = TestInt
	c.Charisma = TestInt
	c.Class = TestString
	c.Constitution = TestInt
	c.Darkvision = true
	c.Dexterity = TestInt
	c.HP = TestInt
	c.History = TestInt
	c.Initiative = TestInt
	c.Insight = TestInt
	c.Intelligence = TestInt
	c.Intimidation = TestInt
	c.Investigation = TestInt
	c.Level = TestInt
	c.Medicine = TestInt
	c.Name = TestString
	c.Nature = TestInt
	c.Perception = TestInt
	c.Performance = TestInt
	c.Persuasion = TestInt
	c.Proficiency = TestInt
	c.Race = TestString
	c.Religion = TestInt
	c.SleightOfHand = TestInt
	c.SpellSlotsOne = TestInt
	c.SpellSlotsTwo = TestInt
	c.SpellSlotsThree = TestInt
	c.SpellSlotsFour = TestInt
	c.SpellSlotsFive = TestInt
	c.SpellSlotsSix = TestInt
	c.SpellSlotsSeven = TestInt
	c.SpellSlotsEight = TestInt
	c.SpellSlotsNine = TestInt
	c.Stealth = TestInt
	c.Strength = TestInt
	c.Survival = TestInt
	c.Subclass = TestString
	c.Subrace = TestString
	c.UsedSpellSlotsOne = TestInt
	c.UsedSpellSlotsTwo = TestInt
	c.UsedSpellSlotsThree = TestInt
	c.UsedSpellSlotsFour = TestInt
	c.UsedSpellSlotsFive = TestInt
	c.UsedSpellSlotsSix = TestInt
	c.UsedSpellSlotsSeven = TestInt
	c.UsedSpellSlotsEight = TestInt
	c.UsedSpellSlotsNine = TestInt
	c.Vision = TestInt
	c.Wisdom = TestInt

	assert.Equal(t, TestInt, c.Acrobatics, "Should be equal")
	assert.Equal(t, TestInt, c.AnimalHandling, "Should be equal")
	assert.Equal(t, TestInt, c.Arcana, "Should be equal")
	assert.Equal(t, TestInt, c.ArmorClass, "Should be equal")
	assert.Equal(t, TestInt, c.Athletics, "Should be equal")
	assert.Equal(t, TestInt, c.Charisma, "Should be equal")
	assert.Equal(t, TestString, c.Class, "Should be equal")
	assert.Equal(t, TestInt, c.Constitution, "Should be equal")
	assert.Equal(t, true, c.Darkvision, "Should be equal")
	assert.Equal(t, TestInt, c.Dexterity, "Should be equal")
	assert.Equal(t, TestInt, c.HP, "Should be equal")
	assert.Equal(t, TestInt, c.History, "Should be equal")
	assert.Equal(t, TestInt, c.Initiative, "Should be equal")
	assert.Equal(t, TestInt, c.Insight, "Should be equal")
	assert.Equal(t, TestInt, c.Intelligence, "Should be equal")
	assert.Equal(t, TestInt, c.Intimidation, "Should be equal")
	assert.Equal(t, TestInt, c.Investigation, "Should be equal")
	assert.Equal(t, TestInt, c.Level, "Should be equal")
	assert.Equal(t, TestInt, c.Medicine, "Should be equal")
	assert.Equal(t, TestString, c.Name, "Should be equal")
	assert.Equal(t, TestInt, c.Nature, "Should be equal")
	assert.Equal(t, TestInt, c.Perception, "Should be equal")
	assert.Equal(t, TestInt, c.Performance, "Should be equal")
	assert.Equal(t, TestInt, c.Persuasion, "Should be equal")
	assert.Equal(t, TestInt, c.Proficiency, "Should be equal")
	assert.Equal(t, TestString, c.Race, "Should be equal")
	assert.Equal(t, TestInt, c.Religion, "Should be equal")
	assert.Equal(t, TestInt, c.SleightOfHand, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsOne, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsTwo, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsThree, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsFour, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsFive, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsSix, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsSeven, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsEight, "Should be equal")
	assert.Equal(t, TestInt, c.SpellSlotsNine, "Should be equal")
	assert.Equal(t, TestInt, c.Stealth, "Should be equal")
	assert.Equal(t, TestInt, c.Strength, "Should be equal")
	assert.Equal(t, TestInt, c.Survival, "Should be equal")
	assert.Equal(t, TestString, c.Subclass, "Should be equal")
	assert.Equal(t, TestString, c.Subrace, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsOne, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsTwo, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsThree, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsFour, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsFive, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsSix, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsSeven, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsEight, "Should be equal")
	assert.Equal(t, TestInt, c.UsedSpellSlotsNine, "Should be equal")
	assert.Equal(t, TestInt, c.Vision, "Should be equal")
	assert.Equal(t, TestInt, c.Wisdom, "Should be equal")
}
