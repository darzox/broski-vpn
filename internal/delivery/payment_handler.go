package delivery

func (s *Delivery) payment(msg Message) error {
	messageString, inlineKeyboard, err := s.usecase.Payment(msg.UserID)
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
