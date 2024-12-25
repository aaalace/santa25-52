package context

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type RequestContext struct {
	Bot      *api.BotAPI
	DbClient *gorm.DB
	Request  *api.Update
}

func BuildBaseContext(bot *api.BotAPI, dbClient *gorm.DB) *RequestContext {
	return &RequestContext{
		Bot:      bot,
		DbClient: dbClient,
		Request:  nil,
	}
}

func (rc *RequestContext) UpdateCurrentContext(request *api.Update) *RequestContext {
	rc.Request = request
	return rc
}
