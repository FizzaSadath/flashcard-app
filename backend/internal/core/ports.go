package core

type CardRepository interface {
	CreateCard(card *Card) error
	GetCardByID(id uint) (*Card, error)
}
