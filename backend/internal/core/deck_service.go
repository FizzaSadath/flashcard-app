package core

import "errors"

type DeckService struct {
	repo DeckRepository
}

func NewDeckService(repo DeckRepository) *DeckService {
	return &DeckService{repo: repo}
}

func (s *DeckService) CreateDeck(userID uint, name string) (*Deck, error) {
	if name == "" {
		return nil, errors.New("deck name cannot be empty")
	}
	deck := &Deck{UserID: userID, Name: name}
	err := s.repo.CreateDeck(deck)
	return deck, err
}

func (s *DeckService) ListDecks(userID uint) ([]Deck, error) {
	return s.repo.ListDecks(userID)
}

func (s *DeckService) DeleteDeck(userID, deckID uint) error {
	deck, err := s.repo.GetDeckByID(deckID)
	if err != nil {
		return err
	}
	if deck.UserID != userID {
		return errors.New("unauthorized")
	}
	return s.repo.DeleteDeck(deckID)
}
