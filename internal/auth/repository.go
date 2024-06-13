package auth

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucas_cda/go-acai-microservices/internal/auth/infra/postgresql"
)

type AuthRepository struct {
	db *pgxpool.Conn
}

func NewAuthRepository(db *pgxpool.Conn) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r AuthRepository) GetUserByEmail(email string) postgresql.User {
	q := postgresql.New(r.db)
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Print("User not found")
	}
	return user
}
