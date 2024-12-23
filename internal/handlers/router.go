package handlers

import (
	"log"
	"santa25-52/internal/context"
)

func HandleRequest(ctx *context.RequestContext) {
	if ctx.Request.Message != nil {
		if ctx.Request.Message.IsCommand() {
			HandleCommand(ctx)
			return
		}
		HandleMessage(ctx)
		return
	} else if ctx.Request.CallbackQuery != nil {
		HandleCallback(ctx)
		return
	}
	log.Println("What the fuck is this user doing")
}
