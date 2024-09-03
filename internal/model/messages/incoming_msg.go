package messages

import "errors"

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type InvoiceSender interface {
	SendInvoice(userId int64, amount int) error
}

type Model struct {
	tgClient   MessageSender
	httpClient InvoiceSender
}

func New(tgClient MessageSender, httpClient InvoiceSender) *Model {
	return &Model{
		tgClient:   tgClient,
		httpClient: httpClient,
	}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {

	switch {
	case msg.Text == "" || msg.UserID == 0:
		return errors.New("cannot send empty message")
	case msg.Text == "/start":
		return s.tgClient.SendMessage("hello", msg.UserID)
	case msg.Text == "/buy":
		return s.httpClient.SendInvoice(msg.UserID, 150)
	case msg.Text == "/paysupport":
		return s.tgClient.SendMessage("Мы напишем вам в лс", msg.UserID)
	default:
		return s.tgClient.SendMessage("the command is unknown", msg.UserID)
	}
}
