package core

// This interface lists all the methods of cards
type CardRepository interface {
	CreateCard(card *Card) error
	GetCardByID(id uint) (*Card, error)
	ListDueCards(limit int) ([]Card, error)
	UpdateCard(card *Card) error
	ListCards() ([]Card, error)
}
