package bot

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"santa25-52/internal/config"
	"santa25-52/internal/context"
	"santa25-52/internal/handlers"
)

type Bot struct {
	Bot     *api.BotAPI
	Updates api.UpdatesChannel
}

func MustLoad(cfg *config.Config) *Bot {
	bot, err := api.NewBotAPI(cfg.BotToken)
	if err != nil {
		panic("Fuck my token")
	}

	u := api.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	return &Bot{
		Bot:     bot,
		Updates: updates,
	}
}

func (b *Bot) SafeLaunch(ctx *context.RequestContext) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Panic occurred | %v\n", r)
				}
			}()

			for update := range b.Updates {
				ctx.UpdateCurrentContext(&update)
				handlers.HandleRequest(ctx)
			}
		}()
	}
}
