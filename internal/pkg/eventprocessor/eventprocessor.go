package eventprocessor

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var startKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("Хочу предложить задачу"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("/help"),
	),
).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")

func ProcessStartComand(bot *telego.Bot, update telego.Update) {

	message := tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Привет. Чем могут быть полезен ?",
	).WithReplyMarkup(startKeyboard)

	_, _ = bot.SendMessage(message)
}

func ProcessHelpComand(bot *telego.Bot, update telego.Update) {

	message := tu.Message(
		tu.ID(update.Message.Chat.ID),
		"На данный момент я умею не так много. \n - Через меня можно предложить задачу на разбор\n - Посмотреть топ 5 желаемых задач от подписчиков.",
	).WithReplyMarkup(startKeyboard)

	_, _ = bot.SendMessage(message)
}

func ProcessNotSupportedComandsComand(bot *telego.Bot, update telego.Update) {
	message := tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Unknown command, use /start",
	).WithReplyMarkup(tu.ReplyKeyboardRemove())

	_, _ = bot.SendMessage(message)
}

func ProcessGetLinkFromReply(bot *telego.Bot, update telego.Update) {
	message := tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Спасибо",
	).WithReplyMarkup(startKeyboard)

	_, _ = bot.SendMessage(message)
}

func ProcessProposeTaskFromMessage(bot *telego.Bot, update telego.Update) {
	message := tu.Message(
		tu.ID(update.Message.Chat.ID),
		"Укажите ссылку на задачу.",
	).WithReplyMarkup(tu.ForceReply())

	_, _ = bot.SendMessage(message)
}
