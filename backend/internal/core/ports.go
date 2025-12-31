package core

// This interface lists all the methods of cards
type CardRepository interface {
	CreateCard(card *Card) error
	GetCardByID(id uint) (*Card, error)
	ListDueCards(userID uint, limit int) ([]Card, error)
	ListCards(userID uint) ([]Card, error)
	UpdateCard(card *Card) error
	GetUserStats(userID uint) (*UserStats, error)
	GetDeckStats(userID uint) ([]DeckStat, error)
	DeleteCard(cardID, userID uint) error
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

type DeckRepository interface {
	CreateDeck(deck *Deck) error
	ListDecks(userID uint) ([]Deck, error)
	DeleteDeck(deckID uint) error
	GetDeckByID(deckID uint) (*Deck, error)
}
