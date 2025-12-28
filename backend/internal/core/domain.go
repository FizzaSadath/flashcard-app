package core

import "time"

type Card struct {
	ID        uint
	UserID    uint
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

type User struct {
	ID       uint
	Email    string
	Username string
	Password string
}

type Deck struct {
	ID        uint
	UserID    uint
	Name      string
	CreatedAt time.Time
}

type Stats struct {
	ID uint
}

type UserStats struct {
	TotalCards    int64
	CardsDue      int64
	NewCards      int64 // Interval = 0
	LearningCards int64 // Interval 1-21
	MatureCards   int64 // Interval > 21
}
