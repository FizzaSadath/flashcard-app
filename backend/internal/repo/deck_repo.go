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
	return nil, nil
}
func (r *PostgresDeckRepo) DeleteDeck(deckID uint) error {
	return nil
}
func (r *PostgresDeckRepo) GetDeckByID(deckID uint) (*core.Deck, error) {
	return nil, nil
}
