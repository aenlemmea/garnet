package data

import "database/sql"

type password struct {
	plaintext *string
	hash      []byte
}

type User struct {
	Id           int      `json:"id"`
	UserId       int      `json:"user_id"`
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	PasswordHash password `json:"-"`
	Preference   string   `json:"preference"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func CreatePostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db: db}
}

type UserStore interface {
	CreateUser(*User) error
	GetUserByUsername(string) (*User, error)
	UpdateUser(*User) error
}

func (pus *PostgresUserStore) CreateUser(user *User) error {
	query := `INSERT INTO Users (user_id, username, email, password_hash, preference)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := pus.db.QueryRow(query, user.UserId, user.Username, user.Email, user.PasswordHash, user.Preference).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (pus *PostgresUserStore) GetUserByUsername(username string) (*User, error) {
	user := &User{
		PasswordHash: password{},
	}

	query := `SELECT id, user_id, username, email, password_hash, preference FROM Users WHERE username = $1`

	err := pus.db.QueryRow(query, username).Scan(
		&user.Id,
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Preference,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (pus *PostgresUserStore) UpdateUser(user *User) error {
	// TODO
	return nil
}
