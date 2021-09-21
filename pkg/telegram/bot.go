package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"tgbot/pkg/config"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	message config.Message
}

func NewBot(bot *tgbotapi.BotAPI, message config.Message) *Bot {
	return &Bot{bot: bot, message: message}
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
		b.HandleMessage(update.Message)

		if update.CallbackQuery != nil {
			b.HandleCallbackDataMenu(update.CallbackQuery)
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

	}
}

func (b *Bot) InitUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.Offset = 0
	u.Limit = 1
	return b.bot.GetUpdatesChan(u)
}
