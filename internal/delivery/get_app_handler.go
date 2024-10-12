package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Delivery) getApp(msg Message) error {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Apple 🍏", "https://itunes.apple.com/app/outline-app/id1356177741"),
			tgbotapi.NewInlineKeyboardButtonURL("Android 🤖", "https://play.google.com/store/apps/details?id=org.outline.android.client"),
		),
	)

	messageString := "Скачать приложение можно используя ссылки:"

	err := s.tgClient.SendMessageWithKeyboard(messageString, msg.UserID, inlineKeyboard)
	if err != nil {
		s.logger.Error("cannot send message:", err)
		return nil
	}

	return nil
}
