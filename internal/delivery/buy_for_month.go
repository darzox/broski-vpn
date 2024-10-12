package delivery

import (
	"database/sql"
	"errors"
)

func (s *Delivery) buyForMonth(msg Message) error {
	_, _, err := s.usecase.GetAccessKey(msg.UserID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if errors.Is(err, sql.ErrNoRows) {
		err = s.usecase.SendInvoiceForMonth(msg.UserID)
		if err != nil {
			return err
		}

		return nil
	}

	messageString, inlineKeyboard, err := s.usecase.BuyForFriendForMonth(msg.UserID)
	if err != nil {
		return err
	}

	err = s.tgClient.SendMessageWithKeyboard(messageString, msg.UserID, *inlineKeyboard)
	if err != nil {
		s.logger.Error("cannot send message:", err)
		return nil
	}

	return nil
}
