package repo

import (
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model

	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (UserEntity) TableName() string {
	return "users"
}

type DeckEntity struct {
	gorm.Model

	UserID uint
	Name   string

	User UserEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (DeckEntity) TableName() string {
	return "decks"
}

type CardEntity struct {
	gorm.Model

	UserID uint       `gorm:"index"`
	User   UserEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	DeckID uint       `gorm:"index"`
	Deck   DeckEntity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Front string `gorm:"type:text;not null"`
	Back  string `gorm:"type:text;not null"`

	Repetitions int     `gorm:"default:0"`
	EaseFactor  float64 `gorm:"default:2.5"`
	Interval    int     `gorm:"default:0"`
}

func (CardEntity) TableName() string {
	return "cards"
}
