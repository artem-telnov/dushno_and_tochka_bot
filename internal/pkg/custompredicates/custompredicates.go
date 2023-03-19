package custompredicates

import "github.com/mymmrac/telego"

func IsNewProposeTask(update telego.Update) bool {
	if update.Message != nil && update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text != "" {
		return update.Message.ReplyToMessage.Text == "Укажите ссылку на задачу."
	}
	return false
}
