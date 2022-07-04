package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreashanson/golang-awake/internal/config"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Have max 10 open database connection at any given time.
const maxOpenDBConn = 10

const maxIdleDBConn = 5

const maxDBLifeTime = 5 * time.Minute

func ConnectSQL(cfg *config.Config) (*sql.DB, error) {
	d, err := NewDatabase(fmt.Sprintf("host=%s port=5432 dbname=%s user=%s password=%s", cfg.Postgres.Host, cfg.Postgres.Name, cfg.Postgres.User, cfg.Postgres.Password))
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetMaxIdleConns(maxIdleDBConn)
	d.SetConnMaxLifetime(maxDBLifeTime)

	if err := testDB(d); err != nil {
		return nil, err
	}

	return d, nil
}

func testDB(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
