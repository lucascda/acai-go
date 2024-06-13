package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func NewPostgresDB(url string) (*postgres, error) {
	var err error
	var pg = &postgres{}

	log.Printf("Connecting to %s", url)
	pg.db, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Error while trying to connect to %s", url)
	}
	log.Print("Connected to Postgres")
	return pg, nil
}

func (p *postgres) GetDB() *pgxpool.Pool {
	return p.db
}

func (p *postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
