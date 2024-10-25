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

func (s *Delivery) buyForMonth(msg Message) error {
	err := s.usecase.BuyForMonth(msg.UserID)
	if err != nil {
		return err

	}

	return nil
}

func (s *Delivery) buyForYear(msg Message) error {
	err := s.usecase.BuyForYear(msg.UserID)
	if err != nil {
		return err

	}

	return nil
}
