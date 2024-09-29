package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/darzox/broski-vpn/internal/clients/http_invoice"
	"github.com/darzox/broski-vpn/internal/clients/outline"
	"github.com/darzox/broski-vpn/internal/repository"
	"github.com/pkg/errors"
)

type Repository interface {
	repository.UserDataStorage
	repository.KeyDataStorage
}

type usecase struct {
	logger               *slog.Logger
	repo                 Repository
	telegramInvoceClient *http_invoice.TelegramHTTPClient
	outlineClient        *outline.OutlineHttpClient
}

func New(logger *slog.Logger, repo Repository, tgInvoiceClinet *http_invoice.TelegramHTTPClient, outlineClient *outline.OutlineHttpClient) *usecase {
	return &usecase{
		logger:               logger,
		repo:                 repo,
		telegramInvoceClient: tgInvoiceClinet,
		outlineClient:        outlineClient,
	}
}

func (u *usecase) Start(chatId int64) (string, error) {
	id, err := u.repo.RegisterUserIfNotExists(context.Background(), chatId)
	if err != nil && err != sql.ErrNoRows {
		return "", errors.Wrap(err, "RegisterUserIfNotExists")
	}

	if id == 0 {
		return "", nil
	}

	accessKey, keyId, err := u.outlineClient.CreateAccessKey()
	if err != nil {
		return "", errors.Wrap(err, "CreateAccessKey")
	}

	expirationDate := time.Now().Add(time.Hour * 24).UTC()

	err = u.repo.CreateUserKey(context.Background(), id, keyId, accessKey, expirationDate)
	if err != nil {
		return "", errors.Wrap(err, "CreateUserKey")
	}

	startMessage := fmt.Sprintf("Добро пожаловать!\n\n👉 Ваш ключ к нашим серверам \n(нажмите на ключ чтобы скопировать его):\n\n`%s`\n\n👉 Тестовый период 24 часа. \n👉 Подписка %d в месяц. \n👉 Для оплаты нажмите /payment\n\n👉 Скачайте приложение и вставьте скопированный ключ\n\n👇👇👇👇👇👇👇👇", accessKey, 100)

	return startMessage, nil
}
