package delivery

func (s *Delivery) createKey(msg Message) error {
	messageString, inlineKeyboard, err := s.usecase.CreateKey(msg.UserID, msg.PaymentInfo)
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
