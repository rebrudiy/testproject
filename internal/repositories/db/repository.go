package db

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"servers/internal/config"
	"servers/internal/models"
)

type PostgresUserRepository struct {
	DB *pgxpool.Pool
}

func NewDB(cfg *config.Config) *PostgresUserRepository {
	db, err := pgxpool.New(context.Background(), cfg.GetDSN())
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	return &PostgresUserRepository{DB: db}
}

func (p *PostgresUserRepository) SaveUser(user models.User) (int, error) {
	slog.Info("Calling save-user")
	slog.Info("Writing user with ID: ", slog.Int("user_id", user.ID))
	query, args, err := sq.Insert("users").
		Columns("id", "email", "first_name", "last_name", "avatar").
		Values(user.ID, user.Email, user.FirstName, user.LastName, user.Avatar).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		slog.Error("Failed to write user", "error", err)
		return 0, err
	}
	_, err = p.DB.Exec(context.Background(), query, args...)
	if err != nil {
		slog.Error("Failed to save user", "error", err)
		return 0, err
	}
	slog.Info("User saved successfully")
	return user.ID, nil
}
