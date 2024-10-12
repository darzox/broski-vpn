package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/darzox/broski-vpn/internal/clients/http_invoice"
	"github.com/darzox/broski-vpn/internal/clients/outline"
	"github.com/darzox/broski-vpn/internal/clients/tg"
	"github.com/darzox/broski-vpn/internal/config"
	"github.com/darzox/broski-vpn/internal/delivery"
	"github.com/darzox/broski-vpn/internal/repository/data_access"
	"github.com/darzox/broski-vpn/internal/usecase"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	config, err := config.New()
	if err != nil {
		logger.Error("config init failed:", err)
		return
	}

	logger.Info("config is loaded : %v", config)

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			config.PostgresDBUserLogin(),
			config.PostgresUserPass(),
			config.PostgresHost(),
			config.PostgresPort(),
			config.PostgresDBName(),
			config.PostgresSslMode()))
	if err != nil {
		logger.Error("migration connection is failed:", "error", err)
		return
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error("migration up is failed:", "error", err)
		return
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.PostgresDBUserLogin(),
		config.PostgresUserPass(),
		config.PostgresHost(),
		config.PostgresPort(),
		config.PostgresDBName(),
		config.PostgresSslMode()))
	if err != nil {
		logger.Error("db connection is failed:", "error", err)
		return
	}

	httpClient, err := http_invoice.NewTelegramHTTPClient(config)
	if err != nil {
		logger.Error("http client init failed:", "error", err)
		return
	}

	httpOutlineClient, err := outline.NewOutlineHttpClient(config)
	if err != nil {
		logger.Error("outline client init failed:", "error", err)
		return
	}
	repo := data_access.NewDb(db)

	usecase := usecase.New(logger, repo, httpClient, httpOutlineClient, config.MonthPriceInXTR())

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("telegram init failed:", "error", err)
	}

	delivery := delivery.New(logger, tgClient, usecase)

	tgClient.ListenUpdates(delivery)
}
