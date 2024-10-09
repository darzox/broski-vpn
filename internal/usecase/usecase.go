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
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type Repository interface {
	repository.UserDataStorage
	repository.KeyDataStorage
	repository.TransactionDataStorage
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

func (u *usecase) Start(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error) {
	id, err := u.repo.RegisterUserIfNotExists(context.Background(), chatId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", nil, errors.Wrap(err, "RegisterUserIfNotExists")
	}

	if id == 0 {
		return "", nil, nil
	}

	accessKey, keyId, err := u.outlineClient.CreateAccessKey()
	if err != nil {
		return "", nil, errors.Wrap(err, "CreateAccessKey")
	}

	expirationDate := time.Now().Add(time.Hour * 24).UTC()

	_, err = u.repo.CreateUserKey(context.Background(), id, keyId, accessKey, expirationDate)
	if err != nil {
		return "", nil, errors.Wrap(err, "CreateUserKey")
	}

	startMessage := fmt.Sprintf("Добро пожаловать!\n\n👉 Ваш ключ к нашим серверам \n(нажмите на ключ чтобы скопировать его):\n\n***`%s`***\n\n👉 Тестовый период 24 часа. \n👉 Подписка %d в месяц. \n👉 Для оплаты нажмите /payment\n\n👉 Скачайте приложение и вставьте скопированный ключ\n\n👇👇👇👇👇👇👇👇", accessKey, 100)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Скачать приложение", "get_app"),
		),
	)

	return startMessage, &inlineKeyboard, nil
}

func (u *usecase) GetUserIdByChatId(chatId int64) (int64, error) {
	id, err := u.repo.GetUserIdByChatId(context.Background(), chatId)
	if err != nil {
		return 0, errors.Wrap(err, "GetUserIdByChatId")
	}

	return id, nil
}

func (u *usecase) GetAccessKey(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error) {
	id, err := u.GetUserIdByChatId(chatId)
	if err != nil {
		return "", nil, errors.Wrap(err, "GetAccessKey.GetUserIdByChatId")
	}

	accessKeys, err := u.repo.GetAccessKeys(context.Background(), id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", nil, errors.Wrap(err, "GetAccessKey.GetAccessKey")
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Скачать приложение", "get_app"),
		),
	)

	var messageString string

	locationTime, _ := time.LoadLocation("Europe/Moscow")

	if len(accessKeys) == 1 {
		expirationDate := accessKeys[0].ExpirationDate.In(locationTime).Format("2006-01-02 15:04:05")
		messageString = fmt.Sprintf("Ваш ключ к нашим серверам \n(нажмите на ключ чтобы скопировать его):\n\n`%s`\n Действителен до: %s(МСК)", accessKeys[0].AccessKeyString, expirationDate)
		return messageString, &inlineKeyboard, nil
	}

	if len(accessKeys) > 1 {
		messageString = "Ваши ключи к нашим серверам:\n\n"
		for _, key := range accessKeys {
			expirationDate := key.ExpirationDate.In(locationTime).Format("2006-01-02 15:04:05")
			messageString += fmt.Sprintf("`%s`\n Действителен до: %s(МСК)\n\n", key.AccessKeyString, expirationDate)
		}

		return messageString, &inlineKeyboard, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		messageString = "Подписка закончилась"
		inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Оплата", "payment"),
			),
		)
	}

	return messageString, &inlineKeyboard, err
}

func (u *usecase) RemoveExpiredKeys(ctx context.Context) error {
	outlineKeyIds, err := u.repo.GetExpiredKeysOutlineIds(ctx)
	if err != nil {
		return errors.Wrap(err, "RemoveExpiredKeys.GetExpiredKeysOutlineIds")
	}

	for _, keyId := range outlineKeyIds {
		err := u.outlineClient.DeleteKey(keyId)
		if err != nil {
			return errors.Wrap(err, "RemoveExpiredKeys.DeleteKey")
		}
	}

	return nil
}

func (u *usecase) SendInvoiceForMonth(chatId int64) error {
	err := u.telegramInvoceClient.SendInvoice(chatId, 1)
	if err != nil {
		return errors.Wrap(err, "SendInvoiceForMonth.SendInvoice")
	}

	return nil
}

func (u *usecase) BuyForFriendForMonth(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error) {
	message := "У вас уже есть ключ доступа\nВы можете купить еще один ключ доступа для друга"

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Купить другу ключ", "payment-for-friend"),
		),
	)

	return message, &inlineKeyboard, nil
}

func (u *usecase) CreateKey(chatId int64, paymentInfo *tgbotapi.SuccessfulPayment) (string, *tgbotapi.InlineKeyboardMarkup, error) {
	id, err := u.GetUserIdByChatId(chatId)
	if err != nil {
		return "", nil, errors.Wrap(err, "CreateKey.GetUserIdByChatId")
	}

	accessKey, keyId, err := u.outlineClient.CreateAccessKey()
	if err != nil {
		return "", nil, errors.Wrap(err, "CreateKey.CreateAccessKey")
	}

	expirationDate := time.Now().Add(time.Hour * 24 * 30).UTC()

	userKeyId, err := u.repo.CreateUserKey(context.Background(), id, keyId, accessKey, expirationDate)
	if err != nil {
		return "", nil, errors.Wrap(err, "CreateKey.CreateUserKey")
	}

	err = u.repo.CreatePaymentTransaction(context.Background(), id, userKeyId, paymentInfo.Currency, paymentInfo.TotalAmount, paymentInfo.InvoicePayload, paymentInfo.TelegramPaymentChargeID, paymentInfo.ProviderPaymentChargeID)
	if err != nil {
		u.logger.Warn("failed to create transaction", err)
	}

	message := fmt.Sprintf("Ваш ключ к нашим серверам \n(нажмите на ключ чтобы скопировать его):\n\n`%s`\n\n👉 Период в 30 дней.\n", accessKey)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Скачать приложение", "get_app"),
		),
	)

	return message, &inlineKeyboard, nil
}
