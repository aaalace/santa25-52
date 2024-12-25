package handlers

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"santa25-52/internal/context"
	"santa25-52/internal/db"
	"santa25-52/internal/ui"
)

func HandleCallback(ctx *context.RequestContext) {
	tgId := ctx.Request.CallbackQuery.Message.Chat.ID
	name := ctx.Request.CallbackQuery.Data
	var member db.Member
	_ = ctx.DbClient.Where("name = ?", name).First(&member)
	_ = ctx.DbClient.Model(&member).Update("tg_id", tgId)

	msg := api.NewMessage(ctx.Request.CallbackQuery.Message.Chat.ID, ui.AskForWishes)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending AskForWishes", err)
	}
}
