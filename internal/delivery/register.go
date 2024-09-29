package delivery

import "log/slog"

type MessageSender interface {
	SendMessage(text string, userID int64) error
	SendAppGetLinks(userID int64) error
}

type InvoiceSender interface {
	SendInvoice(userId int64, amount int) error
}

type Usecase interface {
	Start(userId int64) (string, error)
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
		return s.tgClient.SendAppGetLinks(msg.UserID)
	case msg.Text == "/get_key":
		return s.tgClient.SendMessage("payment", msg.UserID)
	case msg.Text == "/buy_for_month":
		return s.httpClient.SendInvoice(msg.UserID, 150)
	// case msg.Text == "/buy_for_year":
	// 	return s.httpClient.SendInvoice(msg.UserID, 1000)
	case msg.Text == "/support":
		return s.tgClient.SendMessage("Мы напишем вам в лс", msg.UserID)
	default:
		return s.tgClient.SendMessage("the command is unknown", msg.UserID)
	}

	return nil
}
