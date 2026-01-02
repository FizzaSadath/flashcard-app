package repo

import (
	"testing"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=flash_user password=flash_password dbname=flashcard_test_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	err = db.AutoMigrate(&UserEntity{}, &DeckEntity{}, &CardEntity{})
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}

	db.Exec("DELETE FROM cards")
	db.Exec("DELETE FROM decks")
	db.Exec("DELETE FROM users")

	return db
}

func createTestUser(db *gorm.DB) uint {
	user := UserEntity{
		Email:    "test@example.com",
		Username: "test",
		Password: "hashedpassword",
	}
	db.Create(&user)
	return user.ID
}

func createTestDeck(db *gorm.DB, userID uint) uint {
	deck := DeckEntity{
		UserID: userID,
		Name:   "Test Deck",
	}
	db.Create(&deck)
	return deck.ID
}
func TestCreateAndGetCard(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	userID := createTestUser(db)
	deckID := createTestDeck(db, userID)

	originalCard := &core.Card{
		UserID: userID,
		DeckID: deckID,
		Front:  "What is TDD?",
		Back:   "Test Driven Development",
		Stats:  core.InitialStats(),
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
}

func TestListDueCards(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	userID := createTestUser(db)
	deckID := createTestDeck(db, userID)

	card := &core.Card{
		UserID: userID,
		DeckID: deckID,
		Front:  "Due Card",
		Back:   "Answer",
		Stats:  core.InitialStats(),
	}
	repo.CreateCard(card)

	futureCard := &core.Card{
		UserID: userID,
		DeckID: deckID,
		Front:  "Future Card",
		Back:   "Answer",
		Stats:  core.CardStats{Interval: 100, Repetitions: 5, EaseFactor: 2.5},
	}
	repo.CreateCard(futureCard)

	dueCards, err := repo.ListDueCards(userID, deckID, 10)
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

func TestUpdateCard(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	userID := createTestUser(db)
	deckID := createTestDeck(db, userID)

	card := &core.Card{
		UserID: userID,
		DeckID: deckID,
		Front:  "Original Front",
		Back:   "Original Back",
		Stats:  core.InitialStats(),
	}

	if err := repo.CreateCard(card); err != nil {
		t.Fatalf("Failed to create initial card: %v", err)
	}

	card.Front = "Updated Front"
	card.Stats.Repetitions = 5
	card.Stats.Interval = 10

	err := repo.UpdateCard(card)
	if err != nil {
		t.Fatalf("Failed to update card: %v", err)
	}

	updatedCard, err := repo.GetCardByID(card.ID)
	if err != nil {
		t.Fatalf("Failed to get updated card: %v", err)
	}

	if updatedCard.Front != "Updated Front" {
		t.Errorf("Front text did not update. Got %s", updatedCard.Front)
	}
}

func TestGetUserStats(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	userID := createTestUser(db)
	deckID := createTestDeck(db, userID)

	newCard := &core.Card{
		UserID: userID, DeckID: deckID, Front: "New", Back: "A",
		Stats: core.CardStats{Interval: 0},
	}
	repo.CreateCard(newCard)

	learningCard := &core.Card{
		UserID: userID, DeckID: deckID, Front: "Learning", Back: "A",
		Stats: core.CardStats{Interval: 5},
	}
	repo.CreateCard(learningCard)

	matureCard := &core.Card{
		UserID: userID, DeckID: deckID, Front: "Mature", Back: "A",
		Stats: core.CardStats{Interval: 30},
	}
	repo.CreateCard(matureCard)

	db.Exec("UPDATE cards SET updated_at = NOW() - INTERVAL '31 days' WHERE id = ?", matureCard.ID)

	stats, err := repo.GetUserStats(userID)
	if err != nil {
		t.Fatalf("Failed to get stats: %v", err)
	}

	if stats.TotalCards != 3 {
		t.Errorf("Expected Total 3, but got %d", stats.TotalCards)
	}

	if stats.NewCards != 1 {
		t.Errorf("Expected New 1, but got %d", stats.NewCards)
	}

	if stats.LearningCards != 1 {
		t.Errorf("Expected Learning 1, but got %d", stats.LearningCards)
	}

	if stats.MatureCards != 1 {
		t.Errorf("Expected Mature 1, but got %d", stats.MatureCards)
	}

	if stats.CardsDue != 2 {
		t.Errorf("Expected Due 2, but got %d", stats.CardsDue)
	}
}

func TestGetDeckStats(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	userA := createTestUser(db)
	userBEntity := UserEntity{Email: "bob@test.com", Username: "bob", Password: "hash"}
	db.Create(&userBEntity)
	userB := userBEntity.ID

	// Deck 1: "A Mixed Deck" (Will contain various cards)
	deck1 := &core.Deck{UserID: userA, Name: "A Mixed Deck"}
	NewDeckRepo(db).CreateDeck(deck1)

	// Deck 2: "B Empty Deck" (return 0 counts)
	deck2 := &core.Deck{UserID: userA, Name: "B Empty Deck"}
	NewDeckRepo(db).CreateDeck(deck2)

	// Deck 3: "Other User Deck" (Should not appear)
	deck3 := &core.Deck{UserID: userB, Name: "Other Deck"}
	NewDeckRepo(db).CreateDeck(deck3)

	repo.CreateCard(&core.Card{
		UserID: userA, DeckID: deck1.ID, Front: "New", Back: "1",
		Stats: core.CardStats{Interval: 0},
	})

	repo.CreateCard(&core.Card{
		UserID: userA, DeckID: deck1.ID, Front: "Learning", Back: "2",
		Stats: core.CardStats{Interval: 10},
	})

	matureCard := &core.Card{
		UserID: userA, DeckID: deck1.ID, Front: "Mature", Back: "3",
		Stats: core.CardStats{Interval: 30},
	}
	repo.CreateCard(matureCard)
	db.Exec("UPDATE cards SET updated_at = NOW() - INTERVAL '40 days' WHERE id = ?", matureCard.ID)

	stats, err := repo.GetDeckStats(userA)
	if err != nil {
		t.Fatalf("Failed to get deck stats: %v", err)
	}

	if len(stats) != 2 {
		t.Errorf("Expected 2 decks for User A, got %d", len(stats))
	}

	d1 := stats[0]
	if d1.DeckName != "A Mixed Deck" {
		t.Errorf("Expected first deck to be 'A Mixed Deck', got '%s'", d1.DeckName)
	}
	if d1.Total != 3 {
		t.Errorf("Deck1: Expected Total 3, got %d", d1.Total)
	}
	if d1.New != 1 {
		t.Errorf("Deck1: Expected New 1, got %d", d1.New)
	}
	if d1.Learning != 1 { // Interval 10 is between 1-21
		t.Errorf("Deck1: Expected Learning 1, got %d", d1.Learning)
	}
	if d1.Mature != 1 { // Interval 30 is > 21
		t.Errorf("Deck1: Expected Mature 1, got %d", d1.Mature)
	}

	if d1.Due != 2 {
		t.Errorf("Deck1: Expected Due 2, got %d", d1.Due)
	}

	d2 := stats[1]
	if d2.DeckName != "B Empty Deck" {
		t.Errorf("Expected second deck to be 'B Empty Deck', got '%s'", d2.DeckName)
	}
	if d2.Total != 0 {
		t.Errorf("Deck2: Expected Total 0, got %d", d2.Total)
	}
}

func TestDeleteCard(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	userA := createTestUser(db)
	deckA := createTestDeck(db, userA)

	userBEntity := UserEntity{
		Email:    "hacker@test.com",
		Username: "hacker",
		Password: "hashedpassword",
	}
	if err := db.Create(&userBEntity).Error; err != nil {
		t.Fatalf("Failed to create manual user B: %v", err)
	}
	userB := userBEntity.ID

	card := &core.Card{
		UserID: userA,
		DeckID: deckA,
		Front:  "Delete Me",
		Back:   "...",
		Stats:  core.InitialStats(),
	}
	if err := repo.CreateCard(card); err != nil {
		t.Fatalf("Failed to create card: %v", err)
	}

	// User B tries to delete User A's card
	err := repo.DeleteCard(card.ID, userB)
	if err == nil {
		t.Error("Expected error when User B deletes User A's card, got nil")
	}

	// Verify card still exists
	_, err = repo.GetCardByID(card.ID)
	if err != nil {
		t.Error("Card should still exist after unauthorized delete attempt")
	}

	// User A deletes their own card
	err = repo.DeleteCard(card.ID, userA)
	if err != nil {
		t.Fatalf("Failed to delete card: %v", err)
	}

	// Verify card is actually gone
	_, err = repo.GetCardByID(card.ID)
	if err == nil {
		t.Error("Card should be deleted, but GetCardByID found it")
	}

	// Delete Non-Existent Card
	err = repo.DeleteCard(9999, userA)
	if err == nil {
		t.Error("Expected error when deleting non-existent card, got nil")
	}
}
func TestCreateCards_Batch(t *testing.T) {
	db := SetupTestDB(t)
	repo := NewCardRepo(db)

	userID := createTestUser(db)
	deckID := createTestDeck(db, userID)

	var cardsToCreate []core.Card

	for i := 1; i <= 5; i++ {
		frontText := "Front " + string(rune('A'+i))

		cardsToCreate = append(cardsToCreate, core.Card{
			UserID: userID,
			DeckID: deckID,
			Front:  frontText,
			Back:   "Back Value",
			Stats:  core.InitialStats(),
		})
	}

	err := repo.CreateCards(cardsToCreate)
	if err != nil {
		t.Fatalf("CreateCards failed: %v", err)
	}

	savedCards, err := repo.ListCards(userID)
	if err != nil {
		t.Fatalf("Failed to list cards: %v", err)
	}

	if len(savedCards) != 5 {
		t.Errorf("Expected 5 cards to be saved, got %d", len(savedCards))
	}

	firstCard := savedCards[0]

	if firstCard.DeckID != deckID {
		t.Errorf("Card DeckID mismatch. Got %d, want %d", firstCard.DeckID, deckID)
	}
	if firstCard.UserID != userID {
		t.Errorf("Card UserID mismatch. Got %d, want %d", firstCard.UserID, userID)
	}

	if firstCard.Stats.EaseFactor != 2.5 {
		t.Errorf("Card Stats not saved correctly. Got EF %f", firstCard.Stats.EaseFactor)
	}
}
