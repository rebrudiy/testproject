package main

import (
	"log/slog"
	"net/http"
	"os"
	"servers/internal/config"
	"servers/internal/controllers/api"
	"servers/internal/repositories/db"
	"servers/internal/use_cases/createuser"
)

func main() {
	logger := NewLogger()
	logger.Info("App started")
	cfg := config.LoadConfig()

	repository := db.NewDB(cfg)
	defer repository.DB.Close()

	useCase := createuser.NewUserUseCase(repository)

	controller := api.NewController(useCase)

	http.HandleFunc("/user", controller.GetUserHandler)

	logger.Info("Starting server", slog.String("address", ":8080"))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("Failed to start server", slog.String("error", err.Error()))
	}
}

func NewLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
	slog.SetDefault(logger)
	return logger
}
