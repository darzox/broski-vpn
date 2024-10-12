package delivery

func (s *Delivery) start(msg Message) error {
	messageString, inlineKeyboard, err := s.usecase.Start(msg.UserID)
	if err != nil {
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

func (s *Delivery) support(msg Message) error {
	messageString, inlineKeyboard, err := s.usecase.Support(msg.UserID)
	if err != nil {
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
