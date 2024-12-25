package handlers

import (
	"fmt"
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"santa25-52/internal/context"
	"santa25-52/internal/db"
	"santa25-52/internal/game"
	"santa25-52/internal/ui"
)

func HandleCommand(ctx *context.RequestContext) {
	switch ctx.Request.Message.Command() {
	case "start":
		handleStartCommand(ctx)
	case os.Getenv("START_ROUND"):
		handleRoundCommand(ctx)
	default:
		log.Println("What the fuck is this user doing with commands")
	}
}

func handleStartCommand(ctx *context.RequestContext) {
	var members []db.Member
	_ = ctx.DbClient.Find(&members)

	if len(members) == 0 {
		msg := api.NewMessage(ctx.Request.Message.Chat.ID, ui.StartPageEmpty)
		_, err := ctx.Bot.Send(msg)
		if err != nil {
			log.Println("Error sending StartPageEmpty", err)
		}
		return
	}
	msg := api.NewMessage(ctx.Request.Message.Chat.ID, ui.StartPageMessage)
	msg.ReplyMarkup = ui.CreatePeopleListKeyboard(members)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending StartPageMessage", err)
	}
}

func handleRoundCommand(ctx *context.RequestContext) {
	manager := game.Manager{DbClient: ctx.DbClient}
	santaMap := manager.BuildSantaMap()

	for santa, recipient := range santaMap {
		msgText := fmt.Sprintf("Здарова, ты Тайный Санта для %s.\n(S)he wants:\n%s", recipient.Name, recipient.Wish)
		msg := api.NewMessage(santa.TgID, msgText)
		_, err := ctx.Bot.Send(msg)
		if err != nil {
			log.Println(fmt.Sprintf("Error sending SantaResponse to %s", santa.Name), err)
		}

		var member db.Member
		_ = ctx.DbClient.Where("tg_id = ?", recipient.TgID).First(&member)
		if member.FileID != "" {
			photo := api.NewPhoto(santa.TgID, api.FileID(member.FileID))
			_, err := ctx.Bot.Send(photo)
			if err != nil {
				log.Println(fmt.Sprintf("Error sending PhotoResponse to %s", santa.Name), err)
			}
		}
	}
}
