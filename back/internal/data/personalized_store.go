package data

import "database/sql"

type Personalized struct {
	Id            int     `json:"id"`
	UserId        int     `json:"user_id"`
	AggregationId int     `json:"aggregation_id"`
	Score         float64 `json:"score"`
}

type PostgresPersonalizedStore struct {
	db *sql.DB
}

func CreatePostgresPersonalizedStore(db *sql.DB) *PostgresPersonalizedStore {
	return &PostgresPersonalizedStore{db: db}
}

type PersonalizedStore interface {
	GetPersonalizedByUID(uid int64) ([]*Personalized, error)
	PopulatePersonalized(*Personalized) (*Personalized, error)
}

// No communication via API layer. Only to be used internally by the service layer.
func (pp *PostgresPersonalizedStore) PopulatePersonalized(personalized *Personalized) (*Personalized, error) {
	tx, err := pp.db.Begin()
	if err != nil {
		return nil, err
	}
	commit := false
	defer func() {
		if !commit {
			_ = tx.Rollback()
		}
	}()

	query := `INSERT INTO Personalized (user_id, aggregation_id, score) RETURNING id`

	err = tx.QueryRow(query, personalized.UserId, personalized.AggregationId, personalized.Score).Scan(&personalized.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return personalized, nil
}

func (pp *PostgresPersonalizedStore) GetPersonalizedByUID(uid int64) ([]*Personalized, error) {
	personalized := []*Personalized{}
	// TODO
	return personalized, nil
}
