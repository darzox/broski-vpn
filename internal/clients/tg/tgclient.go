package tg

import (
	"log"

	"github.com/darzox/broski-vpn/internal/delivery"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type TokenGetter interface {
	Token() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tokenGetter TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())
	if err != nil {
		return nil, errors.Wrap(err, "NewBotApi")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64) error {
	msg := tgbotapi.NewMessage(userID, text)
	msg.ParseMode = "MarkDown"
	_, err := c.client.Send(msg)
	if err != nil {
		return errors.Wrap(err, "client.Send")
	}

	return nil
}

func (c *Client) SendAppGetLinks(userID int64) error {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Apple 🍏", "https://itunes.apple.com/app/outline-app/id1356177741"),
			tgbotapi.NewInlineKeyboardButtonURL("Android 🤖", "https://play.google.com/store/apps/details?id=org.outline.android.client"),
		),
	)

	msg := tgbotapi.NewMessage(userID, "Скачать приложение можно используя ссылки:")

	msg.ReplyMarkup = inlineKeyboard
	_, err := c.client.Send(msg)
	if err != nil {
		return errors.Wrap(err, "client.SendAppGetLinks")
	}

	return nil
}

func (c *Client) ListenUpdates(router *delivery.Delivery) {
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := router.IncomingMessage(delivery.Message{
				Text:   update.Message.Text,
				UserID: int64(update.Message.From.ID),
			})
			if err != nil {
				log.Println("error proccesing message:", err)
			}
		}
		if update.PreCheckoutQuery != nil {
			log.Printf("userId=%d tries to buy", update.PreCheckoutQuery.From.ID)
		}
	}
}
