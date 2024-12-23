package context

import (
	"github.com/go-redis/redis/v8"
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type RequestContext struct {
	Bot         *api.BotAPI
	CacheClient *redis.Client
	Request     *api.Update
}

func BuildBaseContext(bot *api.BotAPI, cacheClient *redis.Client) *RequestContext {
	return &RequestContext{
		Bot:         bot,
		CacheClient: cacheClient,
		Request:     nil,
	}
}

func (rc *RequestContext) UpdateCurrentContext(request *api.Update) *RequestContext {
	rc.Request = request
	return rc
}
