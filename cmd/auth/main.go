package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lucas_cda/go-acai-microservices/internal/auth"
	"github.com/lucas_cda/go-acai-microservices/pkg/postgres"
)

func main() {
	pg, err := postgres.NewPostgresDB("postgresql://root:example@localhost:5432/db")
	if err != nil {
		panic(err)
	}
	defer pg.Close()
	log.Print("Migrating database...")
	m, err := migrate.New("file://db/migrations", "pgx://root:example@localhost:5432/db")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	api := r.Group("/api/v1/auth")
	api.GET("/health", auth.Health)

	r.Run(":8080")
}
