package handlers

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"santa25-52/internal/context"
	"santa25-52/internal/db"
	"santa25-52/internal/ui"
)

func HandleMessage(ctx *context.RequestContext) {
	if ctx.Request.Message.Photo != nil {
		HandlePhoto(ctx)
	}
	switch ctx.Request.Message.Text {
	default:
		HandleDynamicMessage(ctx)
	}
}

func HandlePhoto(ctx *context.RequestContext) {
	tgId := ctx.Request.Message.Chat.ID

	photo := ctx.Request.Message.Photo[len(ctx.Request.Message.Photo)-1]
	var member db.Member
	_ = ctx.DbClient.Where("tg_id = ?", tgId).First(&member)
	_ = ctx.DbClient.Model(&member).Update("file_id", photo.FileID)

	msg := api.NewMessage(tgId, ui.PhotoAdded)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending PhotoAdded", err)
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
