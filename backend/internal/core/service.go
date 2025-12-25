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

func (s *CardService) CreateCard(front, back string) (*Card, error) {
	newCard := &Card{
		Front: front,
		Back:  back,
		Stats: InitialStats(),
	}

	err := s.repo.CreateCard(newCard)
	if err != nil {
		return nil, err
	}

	return newCard, nil
}

func (s *CardService) ReviewCard(cardID uint, grade int) error {
	// User grades card from 0-5
	if grade < 0 || grade > 5 {
		return errors.New("grade must be between 0 and 5")
	}
	// Fetch card
	card, err := s.repo.GetCardByID(cardID)
	if err != nil {
		return err
	}

	//Update stats
	card.Stats = CalculateReview(card.Stats, grade)

	// Save to DB
	return s.repo.UpdateCard(card)
}

// ListDueCards returns cards that need review
func (s *CardService) ListDueCards(limit int) ([]Card, error) {
	return s.repo.ListDueCards(limit)
}

// ListCards lists all cards, latest first
func (s *CardService) ListCards() ([]Card, error) {
	return s.repo.ListCards()
}
