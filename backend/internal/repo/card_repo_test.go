package repo

import (
	"testing"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {

	dsn := "host=localhost user=flash_user password=flash_password dbname=flashcard_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	err = db.AutoMigrate(&CardEntity{})
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}

	db.Exec("DELETE FROM cards")

	return db
}

func TestCreateAndGetCard(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	originalCard := &core.Card{
		Front: "What is TDD?",
		Back:  "Test Driven Development",
		Stats: core.InitialStats(),
	}

	err := repo.CreateCard(originalCard)
	if err != nil {
		t.Fatalf("Failed to create card: %v", err)
	}

	if originalCard.ID == 0 {
		t.Error("Card ID should not be 0 after saving")
	}

	fetchedCard, err := repo.GetCardByID(originalCard.ID)
	if err != nil {
		t.Fatalf("Failed to get card: %v", err)
	}

	if fetchedCard.Front != "What is TDD?" {
		t.Errorf("Expected 'What is TDD?', got '%s'", fetchedCard.Front)
	}
	if fetchedCard.Stats.EaseFactor != 2.5 {
		t.Errorf("Expected EF 2.5, got %f", fetchedCard.Stats.EaseFactor)
	}
}

func TestListDueCards(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	card := &core.Card{
		Front: "Due Card",
		Back:  "Answer",
		Stats: core.InitialStats(), // Interval = 0
	}
	repo.CreateCard(card)

	futureCard := &core.Card{
		Front: "Future Card",
		Back:  "Answer",
		Stats: core.CardStats{Interval: 100, Repetitions: 5, EaseFactor: 2.5},
	}
	repo.CreateCard(futureCard)

	dueCards, err := repo.ListDueCards(10)
	if err != nil {
		t.Fatalf("Failed to list cards: %v", err)
	}

	if len(dueCards) != 1 {
		t.Errorf("Expected 1 due card, got %d", len(dueCards))
	}

	if len(dueCards) > 0 && dueCards[0].Front != "Due Card" {
		t.Errorf("Expected 'Due Card', got '%s'", dueCards[0].Front)
	}
}
