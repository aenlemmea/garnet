package data

import (
	"database/sql"
)

type Aggregator struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	Blurb      string   `json:"Blurb"`
	Link       string   `json:"link"`
	OriginName string   `json:"origin_name"`
	Tags       []string `json:"tags"`
}

type PostgresAggregatorStore struct {
	db *sql.DB
}

func CreatePostgresAggregatorStore(db *sql.DB) *PostgresAggregatorStore {
	return &PostgresAggregatorStore{db: db}
}

type AggregatorStore interface {
	PopulateAggregator(*Aggregator) (*Aggregator, error)
	GetAggregator(limit, offset int) ([]*Aggregator, error)
	GetAggergatorByID(id int64) (*Aggregator, error)
}

// No communication via API layer. Only to be used internally by the service layer.
func (ps *PostgresAggregatorStore) PopulateAggregator(aggregator *Aggregator) (*Aggregator, error) {
	tx, err := ps.db.Begin()
	if err != nil {
		return nil, err
	}

	commit := false
	defer func() {
		if !commit {
			_ = tx.Rollback()
		}
	}()

	query :=
		`INSERT INTO Aggregation (title, blurb, link, origin_name, tags)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id;
		`

	err = tx.QueryRow(query, aggregator.Title, aggregator.Blurb, aggregator.Link, aggregator.OriginName, aggregator.Tags).Scan(&aggregator.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	commit = true

	return aggregator, nil
}

func (ps *PostgresAggregatorStore) GetAggergatorByID(id int64) (*Aggregator, error) {
	aggregator := &Aggregator{}

	query :=
		`
		SELECT id, title, blurb, link, origin_name
		FROM Aggregation
		WHERE id = $1
	`

	err := ps.db.QueryRow(query, id).Scan(&aggregator.Id, &aggregator.Title, &aggregator.Blurb, &aggregator.Link, &aggregator.OriginName)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return aggregator, nil
}

func (ps *PostgresAggregatorStore) GetAggregator(limit, offset int) ([]*Aggregator, error) {
	aggregators := []*Aggregator{}
	// TODO
	return aggregators, nil
}
