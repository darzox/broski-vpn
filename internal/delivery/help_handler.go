package delivery

func (s *Delivery) help(msg Message) error {
	messageString := `Доступные команды:
/start - Запустить бота
/terms - Посмотреть условия использования
/get_app - Получить ссылку на приложение
/get_key - Получить доступные ключи
/buy_for_month - Купить доступ на месяц
/buy_for_friend_for_month - Купить доступ для друга на месяц
/create_key - Создать новый ключ для VPN
/support - Связаться с поддержкой
`

	err := s.tgClient.SendMessage(messageString, msg.UserID)
	if err != nil {
		s.logger.Error("cannot send message:", err)
		return nil
	}

	return nil
}
