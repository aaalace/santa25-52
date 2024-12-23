package handlers

import (
	"santa25-52/internal/context"
)

func HandleMessage(ctx *context.RequestContext) {
	switch ctx.Request.Message.Text {
	default:
		HandleDynamicMessage(ctx)
	}
}

func HandleDynamicMessage(ctx *context.RequestContext) {

}
