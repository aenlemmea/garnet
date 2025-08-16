package data

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	enverr := godotenv.Load()
	if enverr != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require&channel_binding=require",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)

	db, err := sql.Open("pgx", connStr)

	if err != nil {
		return nil, fmt.Errorf("db: Open %w\n", err)
	}

	fmt.Println("🛢 Connected to database")

	return db, db.Ping()
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("🪿 Migrate: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("🪿 Up: %w", err)
	}

	return nil
}
