package repo

import (
	"gorm.io/gorm"
)

type CardEntity struct {
	gorm.Model

	DeckID uint   `gorm:"index"`
	Front  string `gorm:"type:text;not null"`
	Back   string `gorm:"type:text;not null"`

	Repetitions int     `gorm:"default:0"`
	EaseFactor  float64 `gorm:"default:2.5"`
	Interval    int     `gorm:"default:0"`
}

func (CardEntity) TableName() string {
	return "cards"
}
