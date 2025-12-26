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
func TestListDecks_Isolation(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewDeckRepo(db)

	userA := createTestUser(db)

	userBEntity := UserEntity{Email: "bob@test.com", Password: "hash"}
	db.Create(&userBEntity)
	userB := userBEntity.ID

	repo.CreateDeck(&core.Deck{UserID: userA, Name: "A-1"})
	repo.CreateDeck(&core.Deck{UserID: userA, Name: "A-2"})

	repo.CreateDeck(&core.Deck{UserID: userB, Name: "B-1"})

	decks, err := repo.ListDecks(userA)
	if err != nil {
		t.Fatalf("Failed to list decks: %v", err)
	}

	if len(decks) != 2 {
		t.Errorf("Expected 2 decks for User A, got %d", len(decks))
	}
	if decks[0].Name != "A-1" && decks[0].Name != "A-2" {
		t.Errorf("Unexpected deck name: %s", decks[0].Name)
	}
}

func TestDeleteDeck(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewDeckRepo(db)

	userID := createTestUser(db)
	deck := &core.Deck{UserID: userID, Name: "To Delete"}
	repo.CreateDeck(deck)

	err := repo.DeleteDeck(deck.ID)
	if err != nil {
		t.Fatalf("Failed to delete deck: %v", err)
	}

	decks, _ := repo.ListDecks(userID)
	if len(decks) != 0 {
		t.Errorf("Deck was not deleted, count: %d", len(decks))
	}
}
