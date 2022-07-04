package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreashanson/golang-awake/internal/influencers"
)

type InfluensersRepository struct {
	Connection *sql.DB
}

func NewInfluencersRepo(c *sql.DB) *InfluensersRepository {
	return &InfluensersRepository{Connection: c}
}

func (ir *InfluensersRepository) GetAll() ([]influencers.Influencer, error) {
	allInfluencers := []influencers.Influencer{}

	rows, err := ir.Connection.Query("select * from influencers;")
	if err != nil {
		return nil, err
	}

	var r record
	for rows.Next() {
		if err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Lastname,
			&r.Email,
			&r.CreatedAt,
		); err != nil {
			panic(err)
		}
		allInfluencers = append(allInfluencers, influencers.Influencer{ID: r.ID, Name: r.Name, Lastname: r.Lastname, Email: r.Email, CreatedAt: r.CreatedAt})
	}
	return allInfluencers, nil
}

func (ir InfluensersRepository) GetByID(id string) (influencers.Influencer, error) {
	var r record
	row := ir.Connection.QueryRow(`SELECT * FROM influencers WHERE id = $1`, id)
	if err := row.Scan(&r.ID, &r.Name, &r.Lastname, &r.Email, &r.CreatedAt); err != nil {
		return influencers.Influencer{}, err
	}

	i := influencers.Influencer{ID: r.ID, Name: r.Name, Lastname: r.Lastname, Email: r.Email, CreatedAt: r.CreatedAt}

	fmt.Println(i)

	return i, nil
}

func (ir InfluensersRepository) Create(name, lastname, email string) error {
	_, err := ir.Connection.Exec(`INSERT INTO influencers (name, lastname, email) VALUES ($1, $2, $3)`, name, lastname, email)
	if err != nil {
		return err
	}
	return nil
}

type record struct {
	ID        string
	Name      string
	Lastname  string
	Email     string
	CreatedAt time.Time
}
