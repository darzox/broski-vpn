package delivery

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageSender interface {
	SendMessage(text string, userID int64) error
	SendMessageWithKeyboard(text string, userID int64, keyboard tgbotapi.InlineKeyboardMarkup) error
}

type InvoiceSender interface {
	SendInvoice(userId int64, amount int) error
}

type Usecase interface {
	Start(userId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
	GetAccessKey(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
	SendInvoiceForMonth(chatId int64) error
	BuyForFriendForMonth(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
	CreateKey(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
}

type Delivery struct {
	logger     *slog.Logger
	tgClient   MessageSender
	httpClient InvoiceSender
	usecase    Usecase
}

func New(logger *slog.Logger, tgClient MessageSender, usecase Usecase) *Delivery {
	return &Delivery{
		logger:   logger,
		tgClient: tgClient,
		usecase:  usecase,
	}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Delivery) IncomingMessage(msg Message) error {
	switch {
	case msg.Text == "/start":
		s.start(msg)
	case msg.Text == "/terms":
		return s.tgClient.SendMessage("terms", msg.UserID)
	case msg.Text == "/get_app":
		s.getApp(msg)
	case msg.Text == "/get_key":
		s.getKey(msg)
	case msg.Text == "/buy_for_month":
		s.buyForMonth(msg)
	case msg.Text == "/buy_for_friend_for_month":
		s.buyForFriendForMonth(msg)
	case msg.Text == "/create_key":
		s.createKey(msg)
	case msg.Text == "":
		return nil
	case msg.Text == "/support":
		return s.tgClient.SendMessage("Мы напишем вам в лс", msg.UserID)
	default:
		return s.tgClient.SendMessage("the command is unknown", msg.UserID)
	}

	return nil
}
