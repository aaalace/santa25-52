package handlers

import (
	"fmt"
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"santa25-52/internal/cache"
	"santa25-52/internal/context"
	"santa25-52/internal/ui"
	"strconv"
)

func HandleMessage(ctx *context.RequestContext) {
	switch ctx.Request.Message.Text {
	default:
		HandleDynamicMessage(ctx)
	}
}

func HandleDynamicMessage(ctx *context.RequestContext) {
	wish := ctx.Request.Message.Text
	if wish == "." {
		wish = "Челик не придумал че он(а) хочет"
	}
	ctx.CacheClient.Set(
		ctx.CacheClient.Context(),
		fmt.Sprintf(
			"%s:%s",
			cache.WishesKey,
			strconv.FormatInt(ctx.Request.Message.Chat.ID, 10),
		),
		wish,
		0)

	msg := api.NewMessage(ctx.Request.Message.Chat.ID, ui.WaitingForLastMember)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending WaitingForLastMember", err)
	}
}
