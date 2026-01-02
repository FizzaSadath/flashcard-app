package repo

import (
	"errors"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"gorm.io/gorm"
)

type PostgresCardRepo struct {
	db *gorm.DB
}

func NewCardRepo(db *gorm.DB) core.CardRepository {
	return &PostgresCardRepo{
		db: db,
	}
}

func (r *PostgresCardRepo) CreateCard(card *core.Card) error {
	entity := CardEntity{
		UserID:      card.UserID,
		DeckID:      card.DeckID,
		Front:       card.Front,
		Back:        card.Back,
		Repetitions: card.Stats.Repetitions,
		EaseFactor:  card.Stats.EaseFactor,
		Interval:    card.Stats.Interval,
	}

	result := r.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	card.ID = entity.ID
	card.CreatedAt = entity.CreatedAt

	return nil
}

func (r *PostgresCardRepo) GetCardByID(id uint) (*core.Card, error) {
	var entity CardEntity

	result := r.db.First(&entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("card not found")
		}
		return nil, result.Error
	}
	card := &core.Card{
		ID:     entity.ID,
		UserID: entity.UserID,
		DeckID: entity.DeckID,
		Front:  entity.Front,
		Back:   entity.Back,
		Stats: core.CardStats{
			Repetitions: entity.Repetitions,
			EaseFactor:  entity.EaseFactor,
			Interval:    entity.Interval,
		},
		CreatedAt: entity.CreatedAt,
	}

	return card, nil
}
func (r *PostgresCardRepo) ListDueCards(userID uint, deckID uint, limit int) ([]core.Card, error) {
	var entities []CardEntity

	query := r.db.Where("user_id = ?", userID).
		Where("updated_at::date + make_interval(days => \"interval\"::int) <= CURRENT_DATE")
	if deckID != 0 {
		query = query.Where("deck_id = ?", deckID)
	}

	result := query.Order("updated_at ASC").Limit(limit).Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	var cards []core.Card
	for _, e := range entities {
		cards = append(cards, core.Card{
			ID:     e.ID,
			UserID: e.UserID,
			DeckID: e.DeckID,
			Front:  e.Front,
			Back:   e.Back,
			Stats: core.CardStats{
				Repetitions: e.Repetitions,
				EaseFactor:  e.EaseFactor,
				Interval:    e.Interval,
			},
			CreatedAt: e.CreatedAt,
		})
	}

	return cards, nil
}

func (r *PostgresCardRepo) UpdateCard(card *core.Card) error {

	if card.ID == 0 {
		return errors.New("cannot update card without ID")
	}

	entity := CardEntity{
		DeckID:      card.DeckID,
		UserID:      card.UserID,
		Front:       card.Front,
		Back:        card.Back,
		Repetitions: card.Stats.Repetitions,
		EaseFactor:  card.Stats.EaseFactor,
		Interval:    card.Stats.Interval,
	}
	entity.ID = card.ID
	entity.CreatedAt = card.CreatedAt

	result := r.db.Save(&entity)

	return result.Error
}
func (r *PostgresCardRepo) ListCards(userID uint) ([]core.Card, error) {
	var entities []CardEntity
	result := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}

	var cards []core.Card
	for _, e := range entities {
		cards = append(cards, core.Card{
			ID: e.ID, UserID: e.UserID, DeckID: e.DeckID, Front: e.Front, Back: e.Back,
			Stats:     core.CardStats{Repetitions: e.Repetitions, EaseFactor: e.EaseFactor, Interval: e.Interval},
			CreatedAt: e.CreatedAt,
		})
	}
	return cards, nil
}

func (r *PostgresCardRepo) GetUserStats(userID uint) (*core.UserStats, error) {
	stats := &core.UserStats{}

	// Total Cards
	if err := r.db.Model(&CardEntity{}).Where("user_id = ?", userID).Count(&stats.TotalCards).Error; err != nil {
		return nil, err
	}

	//  Due Cards
	if err := r.db.Model(&CardEntity{}).Where("user_id = ?", userID).
		Where("updated_at::date + make_interval(days => \"interval\"::int) <= CURRENT_DATE").
		Count(&stats.CardsDue).Error; err != nil {
		return nil, err
	}

	// New Cards (Interval == 0)
	r.db.Model(&CardEntity{}).Where("user_id = ? AND \"interval\" = 0", userID).Count(&stats.NewCards)

	// Learning Cards (Interval > 0 AND <= 21)
	r.db.Model(&CardEntity{}).Where("user_id = ? AND \"interval\" > 0 AND \"interval\" <= 21", userID).Count(&stats.LearningCards)

	// Mature Cards (Interval > 21)
	r.db.Model(&CardEntity{}).Where("user_id = ? AND \"interval\" > 21", userID).Count(&stats.MatureCards)

	return stats, nil
}

func (r *PostgresCardRepo) GetDeckStats(userID uint) ([]core.DeckStat, error) {
	var results []core.DeckStat

	query := `
		SELECT 
			d.id as deck_id,
			d.name as deck_name,
			COUNT(c.id) as total,
			SUM(CASE WHEN c.updated_at::date + make_interval(days => c.interval::int) <= CURRENT_DATE THEN 1 ELSE 0 END) as due,
			SUM(CASE WHEN c.interval = 0 THEN 1 ELSE 0 END) as new,
			SUM(CASE WHEN c.interval > 0 AND c.interval <= 21 THEN 1 ELSE 0 END) as learning,
			SUM(CASE WHEN c.interval > 21 THEN 1 ELSE 0 END) as mature
		FROM decks d
		LEFT JOIN cards c ON d.id = c.deck_id
		WHERE d.user_id = ? AND d.deleted_at IS NULL AND c.deleted_at IS NULL
		GROUP BY d.id, d.name
		ORDER BY d.name ASC
	`

	if err := r.db.Raw(query, userID).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (r *PostgresCardRepo) DeleteCard(cardID, userID uint) error {
	result := r.db.Where("id=? AND user_id=?", cardID, userID).Delete(&CardEntity{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Card not found")
	}
	return nil
}

func (r *PostgresCardRepo) CreateCards(cards []core.Card) error {
	var entities []CardEntity

	for _, c := range cards {
		entities = append(entities, CardEntity{
			UserID:      c.UserID,
			DeckID:      c.DeckID,
			Front:       c.Front,
			Back:        c.Back,
			Repetitions: c.Stats.Repetitions,
			EaseFactor:  c.Stats.EaseFactor,
			Interval:    c.Stats.Interval,
		})
	}

	// Insert in batches of 100
	return r.db.CreateInBatches(&entities, 100).Error
}
