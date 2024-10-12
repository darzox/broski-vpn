package delivery

import (
	"database/sql"
	"errors"
)

func (s *Delivery) getKey(msg Message) error {
	messageString, inlineKeyboard, err := s.usecase.GetAccessKey(msg.UserID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.logger.Error("cannot answer to start message:", err)
		return nil
	}

	err = s.tgClient.SendMessageWithKeyboard(messageString, msg.UserID, *inlineKeyboard)
	if err != nil {
		s.logger.Error("cannot send message:", err)
		return nil
	}

	return nil
}
