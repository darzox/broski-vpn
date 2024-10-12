package delivery

func (s *Delivery) buyForFriendForMonth(msg Message) error {
	err := s.usecase.SendInvoiceForMonth(msg.UserID)
	if err != nil {
		return err
	}

	return nil
}
