package repo

import (
	"errors"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"gorm.io/gorm"
)

type PostgresCardRepo struct {
	db *gorm.DB
}

func NewCardRepo(db *gorm.DB) core.CardRepository {
	return &PostgresCardRepo{
		db: db,
	}
}

func (r *PostgresCardRepo) CreateCard(card *core.Card) error {
	entity := CardEntity{
		DeckID:      card.DeckID,
		Front:       card.Front,
		Back:        card.Back,
		Repetitions: card.Stats.Repetitions,
		EaseFactor:  card.Stats.EaseFactor,
		Interval:    card.Stats.Interval,
	}

	result := r.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	card.ID = entity.ID
	card.CreatedAt = entity.CreatedAt

	return nil
}

func (r *PostgresCardRepo) GetCardByID(id uint) (*core.Card, error) {
	var entity CardEntity

	result := r.db.First(&entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("card not found")
		}
		return nil, result.Error
	}
	card := &core.Card{
		ID:     entity.ID,
		DeckID: entity.DeckID,
		Front:  entity.Front,
		Back:   entity.Back,
		Stats: core.CardStats{
			Repetitions: entity.Repetitions,
			EaseFactor:  entity.EaseFactor,
			Interval:    entity.Interval,
		},
		CreatedAt: entity.CreatedAt,
	}

	return card, nil
}
