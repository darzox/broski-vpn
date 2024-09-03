package messages

import (
	"testing"

	mocks "github.com/darzox/telegram-bot.git/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)

	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("hello", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: int64(123),
	})

	assert.NoError(t, err)
}

func Test_OnUknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
	ctrl := gomock.NewController(t)

	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("the command is unknown", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "random command",
		UserID: int64(123),
	})

	assert.NoError(t, err)
}

func Test_EmptyMessageStruct(t *testing.T) {
	model := New(nil)

	err := model.IncomingMessage(Message{})

	assert.EqualError(t, err, "cannot send empty message")
}
