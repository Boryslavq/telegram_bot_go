package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"reflect"
)

const (
	commandStart = "start"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Расходы"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Обратная связь 📲"),
		tgbotapi.NewKeyboardButton("FAQ"),
	),
)

var InlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить расходы", "AddPurchase"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Удалить", "DeletePurchase"),
		tgbotapi.NewInlineKeyboardButtonData("Редактировать", "EditPurchase"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Закрыть окно", "CloseWindow"),
	),
)

var confirmKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Да", "Yes"),
		tgbotapi.NewInlineKeyboardButtonData("Нет", "No"),
	),
)

var cancelKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Отмена", "Cancel"),
	),
)

func (b *Bot) HandleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	switch message.Command() {
	case commandStart:
		msg.Text = b.message.Start
		msg.ReplyMarkup = numericKeyboard
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}

}

func (b *Bot) HandleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	switch message.Text {
	case "Расходы":
		msg.Text = b.message.Expenses
		msg.ReplyMarkup = InlineKeyboard
	case "Обратная связь 📲":
		msg.Text = b.message.Support
		msg.ReplyMarkup = confirmKeyboard
	case "FAQ":
		msg.Text = b.message.FAQ
	}
	_, err := b.bot.Send(msg)
	if err != nil {
		return
	}

}

func (b *Bot) HandleCallbackDataMenu(callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	case "CloseWindow", "No", "Cancel":
		_, err := b.bot.DeleteMessage(tgbotapi.DeleteMessageConfig{ChatID: callback.Message.Chat.ID,
			MessageID: callback.Message.MessageID})
		if err != nil {
			return
		}
	case "AddPurchase":
		msg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, "")
		msg.Text = "Введите запись"
		msg.ReplyMarkup = &cancelKeyboard
		_, err := b.bot.Send(msg)
		if err != nil {
			return
		}
	case "Yes":
		msg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, "")
		msg.Text = b.message.Idea
		b.bot.Send(msg)
	}
}
func (b *Bot) SendMessageToAdmin(message *tgbotapi.Message) {
	if reflect.TypeOf(message.Text).Kind() == reflect.String && message.Text != "" {
		b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Спасибо за обращение!"))
		b.bot.Send(tgbotapi.NewMessage(595259247, message.Text))
	}
}
