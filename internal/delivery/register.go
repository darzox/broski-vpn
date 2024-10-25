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
	BuyForFriend(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
	CreateKey(chatId int64, paymentInfo *tgbotapi.SuccessfulPayment) (string, *tgbotapi.InlineKeyboardMarkup, error)
	Support(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
	Payment(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
	BuyForMonth(chatId int64) (string, *tgbotapi.InlineKeyboardMarkup, error)
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
	Text        string
	UserID      int64
	PaymentInfo *tgbotapi.SuccessfulPayment
}

func (s *Delivery) IncomingMessage(msg Message) error {
	switch {
	case msg.Text == "/start":
		s.start(msg)
	case msg.Text == "/terms":
		return s.tgClient.SendMessage("terms", msg.UserID)
	case msg.Text == "/getapp":
		s.getApp(msg)
	case msg.Text == "/getkey":
		s.getKey(msg)
	case msg.Text == "/buyformonth":
		s.buyForMonth(msg)
	case msg.Text == "/buyforfriendformonth":
		s.buyForFriendForMonth(msg)
	case msg.Text == "/createkey":
		s.createKey(msg)
	case msg.Text == "":
		return nil
	case msg.Text == "/support":
		s.support(msg)
	case msg.Text == "/help":
		s.help(msg)
	case msg.Text == "/instraction":
		s.instraction(msg)
	case msg.Text == "/payment":
		s.payment(msg)
	default:
		s.help(msg)
	}

	return nil
}
