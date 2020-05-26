package storage

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	"go-pg-domain-hex-actor/pkg/listing"
)

// Storage stores beer data
type Storage struct {
	db *sql.DB
}

// NewStorage return a new postgres storage
func NewStorage(postgresURL string) (*Storage, error) {
	cfg, err := pgx.ParseConfig(postgresURL)
	if err != nil {
		return nil, err
	}

	return &Storage{db: stdlib.OpenDB(*cfg)}, nil
}

// GetBeers returns all beers
func (s *Storage) GetBeers() ([]*listing.Beer, error) {
	beers := make([]*listing.Beer, 0)
	rows, err := sq.Select("*").
		From("beers").
		RunWith(s.db).
		Query()
	if err != nil {
		return nil, errors.Errorf("Cannot get beers: %w", err)
	}
	for rows.Next() {
		var beer listing.Beer
		err := rows.Scan(&beer.ID, &beer.Name, &beer.Brewery, &beer.Abv, &beer.ShortDesc, &beer.Created)
		if err != nil {
			return nil, err
		}
		beers = append(beers, &beer)
	}
	return beers, nil
}

// GetBeer returns beer
func (s *Storage) GetBeer(id int) (*listing.Beer, error) {
	sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	beer := listing.Beer{}
	err := sb.Select("*").
		From("beers").
		RunWith(s.db).
		Where("id=?", id).
		QueryRow().
		Scan(&beer.ID, &beer.Name, &beer.Brewery, &beer.Abv, &beer.ShortDesc, &beer.Created)
	if err != nil {
		return nil, errors.Errorf("Cannot get beer: %w", err)
	}
	return &beer, nil
}
