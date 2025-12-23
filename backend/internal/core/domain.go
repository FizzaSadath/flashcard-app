package core

import "time"

type Card struct {
	ID        uint
	DeckID    uint
	Front     string
	Back      string
	Stats     CardStats
	CreatedAt time.Time
}

type CardStats struct {
	Repetitions int
	EaseFactor  float64
	Interval    int
}

func InitialStats() CardStats {
	return CardStats{
		Repetitions: 0,
		EaseFactor:  2.5,
		Interval:    0,
	}
}
