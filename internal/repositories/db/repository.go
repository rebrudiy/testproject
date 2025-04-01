package db

import (
	"log/slog"
	"servers/internal/models"
)

type PostgresUserRepository struct{}

func NewPostgresUserRepository() *PostgresUserRepository {
	return &PostgresUserRepository{}
}

func (*PostgresUserRepository) SaveUser(user models.User) (int, error) {
	slog.Info("Calling save-user")
	slog.Info("Writing user with ID: ", slog.Int("user_id", user.ID))
	if err := DB.Create(&user).Error; err != nil {
		slog.Error("Error to save in db:", err)
		return 0, err
	}
	return user.ID, nil
}
