package core

import "errors"

type CardService struct {
	repo CardRepository
}

func NewCardService(repo CardRepository) *CardService {
	return &CardService{
		repo: repo,
	}
}

// CreateCard now takes userID to link the card to the owner
func (s *CardService) CreateCard(userID uint, front, back string) (*Card, error) {
	newCard := &Card{
		UserID: userID, // Set the owner
		DeckID: 0,
		Front:  front,
		Back:   back,
		Stats:  InitialStats(),
	}

	// We don't need to pass userID separately if it's inside the struct
	err := s.repo.CreateCard(newCard)
	if err != nil {
		return nil, err
	}

	return newCard, nil
}

// ReviewCard now takes userID to ensure you only review YOUR cards
func (s *CardService) ReviewCard(userID uint, cardID uint, grade int) error {
	if grade < 0 || grade > 5 {
		return errors.New("grade must be between 0 and 5")
	}

	// 1. Fetch card
	card, err := s.repo.GetCardByID(cardID)
	if err != nil {
		return err
	}

	// 2. Security Check: Does this card belong to the user?
	if card.UserID != userID {
		return errors.New("unauthorized: this card does not belong to you")
	}

	// 3. Update logic
	card.Stats = CalculateReview(card.Stats, grade)

	return s.repo.UpdateCard(card)
}

// ListDueCards now filters by userID
func (s *CardService) ListDueCards(userID uint, limit int) ([]Card, error) {
	return s.repo.ListDueCards(userID, limit)
}

// ListCards now filters by userID
func (s *CardService) ListCards(userID uint) ([]Card, error) {
	return s.repo.ListCards(userID)
}
