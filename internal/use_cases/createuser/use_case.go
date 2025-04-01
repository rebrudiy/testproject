package createuser

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"servers/internal/models"
)

type Repository interface {
	SaveUser(user models.User) (int, error)
}
type UserUseCase struct {
	repo   Repository
	apiURL string
}

func NewUserUseCase(repo Repository) *UserUseCase {
	return &UserUseCase{repo: repo, apiURL: "https://reqres.in/api/users/2"}
}

func (uc *UserUseCase) CreateUser() (int, error) {
	slog.Info("Calling use-case")
	resp, err := http.Get(uc.apiURL)
	if err != nil {
		slog.Error("Failed to fetch user", "error", err)
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("Failed to close response body", "error", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		err = errors.New("failed to fetch user: non-200 status code")
		slog.Error("Request failed", "status", resp.StatusCode)
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		return 0, err
	}

	var user models.RequestResPayload
	if err := json.Unmarshal(body, &user); err != nil {
		slog.Error("Failed to unmarshal JSON", "error", err)
		return 0, err
	}

	userID, err := uc.repo.SaveUser(user.User)
	if err != nil {
		slog.Error("Failed to save user", "error", err)
		return 0, err
	}

	slog.Info("User created successfully", "user_id", userID)
	return userID, nil
}
