package delivery

func (s *Delivery) start(msg Message) error {
	messageString, err := s.usecase.Start(msg.UserID)
	if err != nil {
		s.logger.Error("cannot answer to start message:", err)
		return nil
	}

	err = s.tgClient.SendMessage(messageString, msg.UserID)
	if err != nil {
		s.logger.Error("cannot send message:", err)
		return nil
	}

	return nil
}
