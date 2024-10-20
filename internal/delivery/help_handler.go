package delivery

func (s *Delivery) help(msg Message) error {
	messageString := `Доступные команды:
/instraction - Инструкция по использованию бота
/get_app - Получить ссылку на приложение
/get_key - Получить доступные ключи
/buy_for_month - Купить доступ на месяц
/support - Связаться с поддержкой
`

	err := s.tgClient.SendMessage(messageString, msg.UserID)
	if err != nil {
		s.logger.Error("cannot send message:", err)
		return nil
	}

	return nil
}
