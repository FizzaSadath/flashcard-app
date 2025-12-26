package repo

import (
	"testing"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
)

func TestCreateDeck(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewDeckRepo(db)

	userID := createTestUser(db)
	deck := &core.Deck{
		UserID: userID,
		Name:   "Test Spanish Deck",
	}

	err := repo.CreateDeck(deck)
	if err != nil {
		t.Fatalf("Failed to create deck: %v", err)
	}

	if deck.ID == 0 {
		t.Error("Deck ID should be set after saving")
	}
}
