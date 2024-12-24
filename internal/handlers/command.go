package handlers

import (
	"fmt"
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"santa25-52/internal/cache"
	"santa25-52/internal/context"
	"santa25-52/internal/game"
	"santa25-52/internal/ui"
	"strconv"
)

func HandleCommand(ctx *context.RequestContext) {
	switch ctx.Request.Message.Command() {
	case "start":
		handleStartCommand(ctx)
	case "round":
		handleRoundCommand(ctx)

	default:
		log.Println("What the fuck is this user doing with commands")
	}
}

func handleStartCommand(ctx *context.RequestContext) {
	names, _ := ctx.CacheClient.LRange(ctx.CacheClient.Context(), cache.TeamsDefaultKey, 0, -1).Result()

	if len(names) == 0 {
		msg := api.NewMessage(ctx.Request.Message.Chat.ID, ui.StartPageEmpty)
		_, err := ctx.Bot.Send(msg)
		if err != nil {
			log.Println("Error sending StartPageEmpty", err)
		}
		return
	}
	msg := api.NewMessage(ctx.Request.Message.Chat.ID, ui.StartPageMessage)
	msg.ReplyMarkup = ui.CreatePeopleListKeyboard(names)
	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("Error sending StartPageMessage", err)
	}
}

func handleRoundCommand(ctx *context.RequestContext) {
	id := strconv.FormatInt(ctx.Request.Message.Chat.ID, 10)
	isMember, _ := ctx.CacheClient.SIsMember(ctx.CacheClient.Context(), cache.AdminSetKey, id).Result()
	if !isMember {
		log.Println(fmt.Sprintf("%v - tried to be admin :)", id))
		return
	}

	manager := game.Manager{
		CacheClient: ctx.CacheClient,
	}
	santaMap := manager.BuildSantaMap()

	for recipientName, santaName := range santaMap {
		recipientId, _ := ctx.CacheClient.Get(
			ctx.CacheClient.Context(),
			fmt.Sprintf("%s:%s", cache.NameToIdKey, recipientName),
		).Result()

		wish, _ := ctx.CacheClient.Get(
			ctx.CacheClient.Context(),
			fmt.Sprintf("%s:%s", cache.WishesKey, recipientId),
		).Result()

		santaId, _ := ctx.CacheClient.Get(
			ctx.CacheClient.Context(),
			fmt.Sprintf("%s:%s", cache.NameToIdKey, santaName),
		).Result()
		santaIdAsInt, _ := strconv.ParseInt(santaId, 10, 64)

		msgText := fmt.Sprintf("Здарова, ты Тайный Санта для %s.\nЕго/Ее пожелания: %s", recipientName, wish)
		msg := api.NewMessage(santaIdAsInt, msgText)
		_, err := ctx.Bot.Send(msg)
		if err != nil {
			log.Println(fmt.Sprintf("Error sending SantaResponse to %s", santaName), err)
		}
	}
}
