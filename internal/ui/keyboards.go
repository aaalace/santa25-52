package ui

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"santa25-52/internal/db"
)

type CallbackButton struct {
	Title string
	Data  string
}

func CreatePeopleListKeyboard(members []db.Member) api.InlineKeyboardMarkup {
	var keyboardRows [][]api.InlineKeyboardButton

	for i := 0; i < len(members); i++ {
		button := api.NewInlineKeyboardButtonData(members[i].Name, members[i].Name)
		keyboardRows = append(keyboardRows, []api.InlineKeyboardButton{button})
	}

	return api.NewInlineKeyboardMarkup(keyboardRows...)
}
