package delivery

func (s *Delivery) help(msg Message) error {
	messageString := `Доступные команды:
/instraction - Инструкция по использованию бота
/getapp - Получить ссылку на приложение
/getkey - Получить доступные ключи
/buyformonth - Купить доступ на месяц
/support - Связаться с поддержкой
`

	err := s.tgClient.SendMessage(messageString, msg.UserID)
	if err != nil {
		s.logger.Error("cannot send message:", err)
		return nil
	}

	return nil
}
