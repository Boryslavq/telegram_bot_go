package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"reflect"
)

const (
	commandStart = "start"
	supportText  = "Напишите свои идеи и предложения!\n" +
		"Поделитесь с нами своими идеями. Мы учтём все пожелания или решим вашу проблему с которой вы столкнулись!"
)

var courseSignMap map[int]*Support

func init() {
	courseSignMap = make(map[int]*Support)
}

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
		msg.Text = "Привет, я ваш личный учётник.\nЯ помогу вам вести контроль ваших растрат."
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
		msg.Text = "У вас пока нет записанных расходов"
		msg.ReplyMarkup = InlineKeyboard
	case "Обратная связь 📲":
		msg.Text = "Вы хотите обратиться в поддержку?"
		msg.ReplyMarkup = confirmKeyboard
	case "FAQ":
		data, err := ioutil.ReadFile("D:/Golang/bot/pkg/telegram/FAQ.txt")
		if err != nil {
			fmt.Println(err)
		}

		msg.Text = string(data)
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
		msg.Text = supportText
		_, err := b.bot.Send(msg)
		if err != nil {
			return
		}
	}
}
func (b *Bot) SendMessageToAdmin(message *tgbotapi.Message) {
	if reflect.TypeOf(message.Text).Kind() == reflect.String && message.Text != "" {
		b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Спасибо за обращение!"))
		b.bot.Send(tgbotapi.NewMessage(595259247, message.Text))
	}
}
