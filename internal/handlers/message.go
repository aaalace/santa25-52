package handlers

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"santa25-52/internal/context"
	"santa25-52/internal/db"
	"santa25-52/internal/ui"
)

func HandleMessage(ctx *context.RequestContext) {
	switch ctx.Request.Message.Text {
	default:
		HandleDynamicMessage(ctx)
	}
}

func HandleDynamicMessage(ctx *context.RequestContext) {
	tgId := ctx.Request.Message.Chat.ID
	wish := ctx.Request.Message.Text
	if wish == "." {
		wish = "Челик не придумал че он хочет"
	}

	var member db.Member
	_ = ctx.DbClient.Where("tg_id = ?", tgId).First(&member)
	_ = ctx.DbClient.Model(&member).Update("wish", wish)

	msg := api.NewMessage(tgId, ui.WaitingForLastMember)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending WaitingForLastMember", err)
	}
}
