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

func HandleCallback(ctx *context.RequestContext) {
	id := strconv.FormatInt(ctx.Request.CallbackQuery.Message.Chat.ID, 10)
	name := ctx.Request.CallbackQuery.Data
	ctx.CacheClient.Set(ctx.CacheClient.Context(), fmt.Sprintf("%s:%s", cache.IdToNameKey, id), name, 0)
	ctx.CacheClient.Set(ctx.CacheClient.Context(), fmt.Sprintf("%s:%s", cache.NameToIdKey, name), id, 0)

	msg := api.NewMessage(ctx.Request.CallbackQuery.Message.Chat.ID, ui.AskForWishes)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending AskForWishes", err)
	}
}
