package core

import "errors"

type CardService struct {
	repo     CardRepository
	deckRepo DeckRepository
}

func NewCardService(repo CardRepository, deckRepo DeckRepository) *CardService {
	return &CardService{
		repo:     repo,
		deckRepo: deckRepo,
	}
}

func (s *CardService) CreateCard(userID uint, deckID uint, front, back string) (*Card, error) {
	newCard := &Card{
		UserID: userID,
		DeckID: deckID,
		Front:  front,
		Back:   back,
		Stats:  InitialStats(),
	}

	err := s.repo.CreateCard(newCard)
	if err != nil {
		return nil, err
	}

	return newCard, nil
}

func (s *CardService) ReviewCard(userID uint, cardID uint, grade int) error {
	if grade < 0 || grade > 5 {
		return errors.New("grade must be between 0 and 5")
	}

	card, err := s.repo.GetCardByID(cardID)
	if err != nil {
		return err
	}

	if card.UserID != userID {
		return errors.New("unauthorized: this card does not belong to you")
	}

	card.Stats = CalculateReview(card.Stats, grade)

	return s.repo.UpdateCard(card)
}

func (s *CardService) ListDueCards(userID uint, limit int) ([]Card, error) {
	return s.repo.ListDueCards(userID, limit)
}

func (s *CardService) ListCards(userID uint) ([]Card, error) {
	return s.repo.ListCards(userID)
}
