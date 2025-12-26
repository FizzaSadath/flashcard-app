package repo

import (
	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"gorm.io/gorm"
)

type PostgresDeckRepo struct {
	db *gorm.DB
}

func NewDeckRepo(db *gorm.DB) core.DeckRepository {
	return &PostgresDeckRepo{db: db}
}

func (r *PostgresDeckRepo) CreateDeck(deck *core.Deck) error {
	entity := DeckEntity{
		UserID: deck.UserID,
		Name:   deck.Name,
	}

	result := r.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	deck.ID = entity.ID
	deck.CreatedAt = entity.CreatedAt
	return nil
}

func (r *PostgresDeckRepo) ListDecks(userID uint) ([]core.Deck, error) {
	var entities []DeckEntity

	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&entities).Error
	if err != nil {
		return nil, err
	}

	var decks []core.Deck
	for _, e := range entities {
		decks = append(decks, core.Deck{
			ID:        e.ID,
			UserID:    e.UserID,
			Name:      e.Name,
			CreatedAt: e.CreatedAt,
		})
	}
	return decks, nil
}
func (r *PostgresDeckRepo) DeleteDeck(deckID uint) error {
	return r.db.Delete(&DeckEntity{}, deckID).Error
}

func (r *PostgresDeckRepo) GetDeckByID(deckID uint) (*core.Deck, error) {
	var entity DeckEntity
	if err := r.db.First(&entity, deckID).Error; err != nil {
		return nil, err
	}
	return &core.Deck{ID: entity.ID, UserID: entity.UserID, Name: entity.Name}, nil
}
