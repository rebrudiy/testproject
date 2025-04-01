package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type UseCase interface {
	CreateUser() (int, error)
}

type Controller struct {
	useCase UseCase
}

func NewController(useCase UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (c *Controller) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received request", slog.String("endpoint", "/user"))

	userID, err := c.useCase.CreateUser()
	if err != nil {
		slog.Error("Failed to create user", slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Find user %d", userID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

	slog.Info("User successfully created", slog.Int("user_id", userID))
}
