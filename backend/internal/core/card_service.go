package core

import (
	"encoding/csv"
	"errors"
	"io"
)

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

func (s *CardService) GetStats(userID uint) (*UserStats, error) {
	return s.repo.GetUserStats(userID)
}

func (s *CardService) GetDeckStats(userID uint) ([]DeckStat, error) {
	return s.repo.GetDeckStats(userID)
}
func (s *CardService) DeleteCard(userID uint, cardID uint) error {
	return s.repo.DeleteCard(cardID, userID)
}

func (s *CardService) ImportCards(userID uint, deckID uint, file io.Reader) (int, error) {
	deck, err := s.deckRepo.GetDeckByID(deckID)
	if err != nil {
		return 0, err
	}
	if deck.UserID != userID {
		return 0, errors.New("unauthorized")
	}

	reader := csv.NewReader(file)
	reader.LazyQuotes = true // Helps with messy quotes

	var cards []Card
	lineCount := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		if lineCount == 0 && (record[0] == "Front" || record[0] == "front") {
			lineCount++
			continue
		}

		if len(record) < 2 {
			continue
		}

		cards = append(cards, Card{
			UserID: userID,
			DeckID: deckID,
			Front:  record[0],
			Back:   record[1],
			Stats:  InitialStats(),
		})
		lineCount++
	}

	if len(cards) == 0 {
		return 0, errors.New("no valid cards found in CSV")
	}

	if err := s.repo.CreateCards(cards); err != nil {
		return 0, err
	}

	return len(cards), nil
}
