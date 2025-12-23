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
func (r *PostgresCardRepo) ListDueCards(limit int) ([]core.Card, error) {
	var entities []CardEntity

	// 1. The Query
	// We use raw SQL syntax for the date math because it's cleaner than GORM maps for this.
	// Logic: "Select * FROM cards WHERE updated_at + (interval days) < NOW()"
	// 'interval' is a keyword in SQL, so we quote our column name "interval"
	query := r.db.Where("updated_at + make_interval(days => \"interval\"::int) <= NOW()")

	// Apply Limit and Order (Oldest due first)
	result := query.Order("updated_at ASC").Limit(limit).Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	// 2. Map Entities -> Domain Models
	var cards []core.Card
	for _, e := range entities {
		cards = append(cards, core.Card{
			ID:     e.ID,
			DeckID: e.DeckID,
			Front:  e.Front,
			Back:   e.Back,
			Stats: core.CardStats{
				Repetitions: e.Repetitions,
				EaseFactor:  e.EaseFactor,
				Interval:    e.Interval,
			},
			CreatedAt: e.CreatedAt,
		})
	}

	return cards, nil
}

func (r *PostgresCardRepo) UpdateCard(card *core.Card) error {
	return nil
}
