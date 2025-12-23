package core

type CardRepository interface {
	CreateCard(card *Card) error
	GetCardByID(id uint) (*Card, error)
	ListDueCards(limit int) ([]Card, error)
	UpdateCard(card *Card) error
}
