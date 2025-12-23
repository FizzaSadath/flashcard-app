package core

type CardRepository interface {
	CreateCard(card *Card)
	GetCardByID(id uint) (*Card, error)
}
