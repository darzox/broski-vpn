package main

import (
	"fmt"
	"log"

	"github.com/darzox/telegram-bot.git/internal/clients/http_invoice"
	"github.com/darzox/telegram-bot.git/internal/clients/tg"
	"github.com/darzox/telegram-bot.git/internal/config"
	model "github.com/darzox/telegram-bot.git/internal/model/messages"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	m, err := migrate.New(
		"../../migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			config.PostgresDBUserLogin(),
			config.PostgresUserPass(),
			config.PostgresHost(),
			config.PostgresPort(),
			config.PostgresSslMode(),
			config.PostgresDBName()))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.PostgresDBUserLogin(),
		config.PostgresUserPass(),
		config.PostgresHost(),
		config.PostgresPort(),
		config.PostgresSslMode(),
		config.PostgresDBName()))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("telegram init failed:", err)
	}

	httpClient, err := http_invoice.NewTelegramHTTPClient(config)

	model := model.New(tgClient, httpClient)

	tgClient.ListenUpdates(model)
}
