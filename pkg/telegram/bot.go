package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates, err := b.InitUpdatesChannel()
	if err != nil {
		return nil
	}
	b.HandleUpdates(updates)

	return nil
}

func (b *Bot) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	notification := tgbotapi.NewMessage(595259247, "Бот запущен")
	b.bot.Send(notification)
	for update := range updates {
		if update.CallbackQuery != nil {
			b.HandleCallbackDataMenu(update.CallbackQuery)
		}
		if update.CallbackQuery != nil && update.CallbackQuery.Data == "Yes" {
			b.SendMessageToAdmin(update.Message)
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			err := b.HandleCommand(update.Message)
			if err != nil {
				return
			}
			continue
		}
		b.HandleMessage(update.Message)

	}
}

func (b *Bot) InitUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.Offset = 0
	u.Limit = 1
	return b.bot.GetUpdatesChan(u)
}
